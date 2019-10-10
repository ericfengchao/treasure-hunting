package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"sync"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service/models"
)

type svc struct {
	role           models.Role
	playerId       string
	registry       *game_pb.Registry
	gridSize       int
	treasureAmount int

	rwLock *sync.RWMutex

	game     models.Gamer
	gameCopy *game_pb.CopyRequest

	masterConn   *grpc.ClientConn
	masterNode   *game_pb.Player
	masterClient game_pb.GameServiceClient

	slaveConn   *grpc.ClientConn
	slaveNode   *game_pb.Player
	slaveClient game_pb.GameServiceClient
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.role == models.PrimaryNode {
		fmt.Fprint(w, s.game.GetGridView())
	} else if s.role == models.BackupNode {
		backupView := &models.ViewableGameStats{Grid: s.gameCopy.GetGrid()}
		fmt.Fprint(w, backupView.GetGridView())
	}
}

func (s *svc) RequestCopy(ctx context.Context, req *game_pb.RequestCopyRequest) (*game_pb.RequestCopyResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	if s.role == models.PrimaryNode {
		if s.game != nil {
			return &game_pb.RequestCopyResponse{
				Status: game_pb.RequestCopyResponse_OK,
				Copy:   s.game.GetSerialisedGameStats(),
			}, nil
		} else {
			log.Printf("player %s requested gamecopy from player %s, but game is nil", req.GetPlayerId(), s.playerId)
			return &game_pb.RequestCopyResponse{
				Status: game_pb.RequestCopyResponse_NULL_ERROR,
			}, nil
		}
	} else if s.role == models.BackupNode {
		return &game_pb.RequestCopyResponse{
			Status: game_pb.RequestCopyResponse_OK,
			Copy:   s.gameCopy,
		}, nil
	} else {
		return &game_pb.RequestCopyResponse{
			Status: game_pb.RequestCopyResponse_I_AM_NOT_PRIMARY,
		}, nil
	}
}

func (s *svc) StatusCopy(ctx context.Context, req *game_pb.CopyRequest) (*game_pb.CopyResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	log.Printf("player-%s received game copy req version %d", s.playerId, req.GetStateVersion())
	if s.gameCopy.GetStateVersion() <= req.GetStateVersion() {
		fmt.Printf("player-%s updating game copy from %d to %d\n", s.playerId, s.gameCopy.GetStateVersion(), req.GetStateVersion())
		//fmt.Println("current", s.gameCopy)
		//fmt.Println("new", req)
		s.gameCopy = req
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_OK,
		}, nil
	} else {
		log.Printf("received version is old! current: %d, received: %d", s.gameCopy.GetStateVersion(), req.GetStateVersion())
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_NULL_ERROR,
		}, nil
	}
}

func (s *svc) MovePlayer(ctx context.Context, req *game_pb.MoveRequest) (*game_pb.MoveResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	log.Println(req)
	// only take requests when i am a primary server node
	// only take requests when i am a primary server node
	if s.role == models.BackupNode {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_I_AM_ONLY_BACKUP,
		}, nil
	} else if s.role == models.PlayerNode {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_I_AM_NOT_A_SERVER,
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
			return &game_pb.MoveResponse{
				Status: game_pb.MoveResponse_SLAVE_INIT_IN_PROGRESS,
			}, nil
		}
		syncResp, syncErr := s.slaveClient.StatusCopy(context.Background(), s.game.GetSerialisedGameStats())
		if syncErr != nil || syncResp.GetStatus() != game_pb.CopyResponse_OK {
			log.Printf("syncing with slave %s not successful: %s, %v", s.slaveNode.GetPlayerId(), syncResp.GetStatus().String(), syncErr)
			// refresh slave
			s.roleSetup()
			return &game_pb.MoveResponse{
				Status: game_pb.MoveResponse_SLAVE_INIT_IN_PROGRESS,
			}, nil
		}
	}

	err := s.game.MovePlayer(req.GetId(), models.Movement(int(req.GetMove())))
	if err == models.InvalidCoordinates {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_INVALID_INPUT,
		}, nil
	} else if err == models.PlaceAlreadyTaken {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_SLOT_TAKEN,
		}, nil
	} else if err == models.SlaveIsDown {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_SLAVE_INIT_IN_PROGRESS,
		}, nil
	} else if err != nil {
		// unknown error occurred. e.g. io timeout. caller should retry
		return nil, err
	}

	// real sync. if failure happens, the next request will detect
	if len(s.registry.GetPlayerList()) > 1 && s.slaveClient != nil {
		fmt.Println(s.slaveClient.StatusCopy(context.Background(), s.game.GetSerialisedGameStats()))
	}

	return &game_pb.MoveResponse{
		PlayerStates: s.game.GetGameStates(),
		Status:       game_pb.MoveResponse_OK,
	}, nil
}

func (s *svc) Heartbeat(ctx context.Context, req *game_pb.HeartbeatRequest) (*game_pb.HeartbeatResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	// if registry changes, there might be a role change as well
	//log.Println(fmt.Sprintf("player %s received heartbeat from %s, registry version: %d", s.playerId, req.PlayerId, req.Registry.GetVersion()))
	if s.registry.GetVersion() < req.GetRegistry().GetVersion() {
		s.registry = req.GetRegistry()
		log.Println("new registry", s.registry)
		s.roleSetup()
	}

	return &game_pb.HeartbeatResponse{}, nil
}

func (s *svc) deriveRole() models.Role {
	if GetPrimaryServer(s.registry).GetPlayerId() == s.playerId {
		return models.PrimaryNode
	} else if GetBackupServer(s.registry).GetPlayerId() == s.playerId {
		return models.BackupNode
	} else {
		return models.PlayerNode
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
	case models.BackupNode:
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
				resp, copyErr := s.masterClient.RequestCopy(context.Background(), &game_pb.RequestCopyRequest{
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
	case models.PrimaryNode:
		// sync slave
		if oldRole == models.BackupNode {
			s.game = models.NewGameFromGameCopy(s.gameCopy)
		} else if oldRole != models.PrimaryNode {
			s.game = models.NewGame(s.gridSize, s.treasureAmount)
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

func (s *svc) UpdateLocalRegistry(registry *game_pb.Registry) {
	s.rwLock.Lock()
	s.rwLock.Unlock()

	s.registry = registry
	s.roleSetup()
}

func (s *svc) GetLocalRegistry() *game_pb.Registry {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	return s.registry
}

func NewGameSvc(playerId string, gridSize int, treasureAmount int, registry *game_pb.Registry) GameService {
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
