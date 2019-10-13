package treasure_hunting

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"sync"
)

type GameService interface {
	GameServiceServer
	http.Handler
	GetLocalRegistry() *Registry
	SyncPlayerStates(resp *MoveResponse)
	UpdateLocalRegistry(*Registry)
}

type svc struct {
	role           Role
	playerId       string
	registry       *Registry
	gridSize       int
	treasureAmount int

	playerStatesFromServer []*PlayerState

	rwLock *sync.RWMutex

	game     Gamer
	gameCopy *CopyRequest

	masterConn   *grpc.ClientConn
	masterNode   *Player
	masterClient GameServiceClient

	slaveConn   *grpc.ClientConn
	slaveNode   *Player
	slaveClient GameServiceClient
}

func (s *svc) SyncPlayerStates(resp *MoveResponse) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.playerStatesFromServer = resp.GetPlayerStates()
	if s.gameCopy.GetStateVersion() <= resp.GetCopy().GetStateVersion() {
		s.gameCopy = resp.GetCopy()
	}
	return
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	primaryNode := GetPrimaryServer(s.registry)
	backupNode := GetBackupServer(s.registry)
	gameViewComponents := PlayerStatesView{
		SelfId:   s.playerId,
		MasterId: primaryNode.GetPlayerId(),
		SlaveId:  backupNode.GetPlayerId(),
	}
	if s.role == PrimaryNode {
		gameViewComponents.PlayerStates = s.game.GetSerialisedGameStats().GetPlayerStates()
		masterView := &BackupViewGameStates{
			PlayerStatesView: gameViewComponents,
			Grid:             s.game.GetSerialisedGameStats().GetGrid(),
		}
		fmt.Fprint(w, masterView.GetViews())
	} else if s.role == BackupNode {
		gameViewComponents.PlayerStates = s.gameCopy.GetPlayerStates()
		backupView := &BackupViewGameStates{
			PlayerStatesView: gameViewComponents,
			Grid:             s.gameCopy.GetGrid(),
		}
		fmt.Fprint(w, backupView.GetViews())
	} else if s.role == PlayerNode {
		gameViewComponents.PlayerStates = s.gameCopy.GetPlayerStates()
		ps := BackupViewGameStates{
			Grid:             s.gameCopy.GetGrid(),
			PlayerStatesView: gameViewComponents,
		}
		fmt.Fprint(w, ps.GetViews())
	}
}

func (s *svc) RequestCopy(ctx context.Context, req *RequestCopyRequest) (*RequestCopyResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	if s.role == PrimaryNode {
		if s.game != nil {
			return &RequestCopyResponse{
				Status: RequestCopyResponse_OK,
				Copy:   s.game.GetSerialisedGameStats(),
			}, nil
		} else {
			log.Printf("player %s requested gamecopy from player %s, but game is nil", req.GetPlayerId(), s.playerId)
			return &RequestCopyResponse{
				Status: RequestCopyResponse_NULL_ERROR,
			}, nil
		}
	} else if s.role == BackupNode {
		return &RequestCopyResponse{
			Status: RequestCopyResponse_OK,
			Copy:   s.gameCopy,
		}, nil
	} else {
		return &RequestCopyResponse{
			Status: RequestCopyResponse_I_AM_NOT_PRIMARY,
		}, nil
	}
}

func (s *svc) StatusCopy(ctx context.Context, req *CopyRequest) (*CopyResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	log.Printf("player-%s received game copy req version %d", s.playerId, req.GetStateVersion())
	if s.gameCopy.GetStateVersion() <= req.GetStateVersion() {
		fmt.Printf("player-%s updating game copy from %d to %d\n", s.playerId, s.gameCopy.GetStateVersion(), req.GetStateVersion())
		//fmt.Println("current", s.gameCopy)
		//fmt.Println("new", req)
		s.gameCopy = req
		return &CopyResponse{
			Status: CopyResponse_OK,
		}, nil
	} else {
		log.Printf("received version is old! current: %d, received: %d", s.gameCopy.GetStateVersion(), req.GetStateVersion())
		return &CopyResponse{
			Status: CopyResponse_NULL_ERROR,
		}, nil
	}
}

func (s *svc) MovePlayer(ctx context.Context, req *MoveRequest) (*MoveResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	log.Println(req)
	// only take requests when i am a primary server node
	// only take requests when i am a primary server node
	if s.role == BackupNode {
		return &MoveResponse{
			Status: MoveResponse_I_AM_ONLY_BACKUP,
		}, nil
	} else if s.role == PlayerNode {
		return &MoveResponse{
			Status: MoveResponse_I_AM_NOT_A_SERVER,
		}, nil
	}

	// only master will deal with TakeSlot requests
	// 1. update locally
	// 2. serialise game status and sync with slave
	// 3. reply player

	// slave healthcheck
	if len(s.registry.GetPlayerList()) > 1 {
		if s.slaveClient == nil {
			s.roleSetup()
			return &MoveResponse{
				Status: MoveResponse_SLAVE_INIT_IN_PROGRESS,
			}, nil
		}
		syncResp, syncErr := s.slaveClient.StatusCopy(context.Background(), s.game.GetSerialisedGameStats())
		if syncErr != nil || syncResp.GetStatus() != CopyResponse_OK {
			log.Printf("syncing with slave %s not successful: %s, %v", s.slaveNode.GetPlayerId(), syncResp.GetStatus().String(), syncErr)
			// refresh slave
			s.roleSetup()
			return &MoveResponse{
				Status: MoveResponse_SLAVE_INIT_IN_PROGRESS,
			}, nil
		}
	}

	err := s.game.MovePlayer(req.GetId(), Movement(int(req.GetMove())))
	gs := s.game.GetGameStates()
	if err == InvalidCoordinates {
		return &MoveResponse{
			Status:       MoveResponse_INVALID_INPUT,
			PlayerStates: gs,
		}, nil
	} else if err == PlaceAlreadyTaken {
		return &MoveResponse{
			Status:       MoveResponse_SLOT_TAKEN,
			PlayerStates: gs,
		}, nil
	} else if err == SlaveIsDown {
		return &MoveResponse{
			Status:       MoveResponse_SLAVE_INIT_IN_PROGRESS,
			PlayerStates: gs,
		}, nil
	} else if err != nil {
		// unknown error occurred. e.g. io timeout. caller should retry
		return nil, err
	}

	// real sync. if failure happens, the next request will detect
	if len(s.registry.GetPlayerList()) > 1 && s.slaveClient != nil {
		fmt.Println(s.slaveClient.StatusCopy(context.Background(), s.game.GetSerialisedGameStats()))
	}

	return &MoveResponse{
		PlayerStates: gs,
		Status:       MoveResponse_OK,
		Copy:         s.game.GetSerialisedGameStats(),
	}, nil
}

func (s *svc) Heartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	// if registry changes, there might be a role change as well
	//log.Println(fmt.Sprintf("player %s received heartbeat from %s, registry version: %d", s.playerId, req.PlayerId, req.Registry.GetVersion()))
	if s.registry.GetVersion() < req.GetRegistry().GetVersion() {
		s.registry = req.GetRegistry()
		log.Println("new registry", s.registry)
		s.roleSetup()
	}

	return &HeartbeatResponse{}, nil
}

func (s *svc) deriveRole() Role {
	if GetPrimaryServer(s.registry).GetPlayerId() == s.playerId {
		return PrimaryNode
	} else if GetBackupServer(s.registry).GetPlayerId() == s.playerId {
		return BackupNode
	} else {
		return PlayerNode
	}
}

// derive role, if there's any role change, setup the svc based on the role
func (s *svc) roleSetup() {
	//s.rwLock.Lock()
	//defer s.rwLock.Unlock()

	log.Println("player", s.playerId, s.registry)
	oldRole := s.role
	s.role = s.deriveRole()

	if oldRole != s.role {
		log.Printf("player %s role change from %s to %s", s.playerId, oldRole, s.role)
	}

	switch s.role {
	case BackupNode:
		primaryNode := GetPrimaryServer(s.registry)
		log.Printf("slave player-%s refreshing primary player-%s, prev : %s", s.playerId, primaryNode.GetPlayerId(), s.masterNode.GetPlayerId())
		if s.masterNode.GetPlayerId() != primaryNode.GetPlayerId() {
			s.masterNode = primaryNode
			if s.masterConn != nil {
				s.masterConn.Close()
			}
			s.masterConn, s.masterClient = ConnectToPlayer(primaryNode)
			if s.masterClient != nil {
				fmt.Printf("player %s requesting game copy from master %s\n", s.playerId, s.masterNode.GetPlayerId())
				resp, copyErr := s.masterClient.RequestCopy(context.Background(), &RequestCopyRequest{
					PlayerId: s.playerId,
				})
				if copyErr != nil {
					fmt.Printf("player %s requesting game copy from master %s failed: %s\n", s.playerId, s.masterNode.GetPlayerId(), copyErr.Error())
				} else {
					if s.gameCopy.GetStateVersion() <= resp.GetCopy().GetStateVersion() {
						s.gameCopy = resp.Copy
						fmt.Println("successfully get game copy", s.gameCopy)
					}
				}
			}
		}
	case PrimaryNode:
		// sync slave
		if oldRole == BackupNode {
			s.game = NewGameFromGameCopy(s.gameCopy)
		} else if oldRole != PrimaryNode {
			s.game = NewGame(s.gridSize, s.treasureAmount)
		}
		s.game.CleanupPlayer(s.registry.GetPlayerList())
		s.gameCopy = s.game.GetSerialisedGameStats()
		if len(s.registry.GetPlayerList()) <= 1 {
			return
		}
		newBackupNode := GetBackupServer(s.registry)
		if s.slaveNode.GetPlayerId() != newBackupNode.GetPlayerId() {
			if s.slaveConn != nil {
				s.slaveConn.Close()
			}
			s.slaveNode = newBackupNode
			s.slaveConn, s.slaveClient = ConnectToPlayer(newBackupNode)
			if s.slaveClient != nil {
				log.Println("new slave first sync", s.slaveNode)
				resp, syncErr := s.slaveClient.StatusCopy(context.Background(), s.game.GetSerialisedGameStats())
				log.Println(resp, syncErr)
			}
		}
	}
	return
}

func (s *svc) UpdateLocalRegistry(registry *Registry) {
	s.rwLock.Lock()
	s.rwLock.Unlock()

	s.registry = registry
	s.roleSetup()
}

func (s *svc) GetLocalRegistry() *Registry {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	return s.registry
}

func NewGameSvc(playerId string, gridSize int, treasureAmount int, registry *Registry) GameService {
	s := &svc{
		playerId:       playerId,
		registry:       registry,
		gridSize:       gridSize,
		treasureAmount: treasureAmount,
		rwLock:         &sync.RWMutex{},
	}
	s.roleSetup()
	return s
}
