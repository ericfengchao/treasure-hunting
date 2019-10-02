package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service/models"
)

type svc struct {
	role     models.Role
	playerId string
	registry *game_pb.Registry

	rwLock *sync.RWMutex

	game     models.Gamer
	gameCopy *game_pb.CopyRequest
	slave    game_pb.GameServiceClient
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.role == models.PrimaryNode {
		fmt.Fprint(w, s.game.GetGridView())
	} else if s.role == models.BackupNode {
		backupView := &models.ViewableGameStats{Grid: s.gameCopy.GetGrid()}
		fmt.Fprint(w, backupView.GetGridView())
	}
}

func (s *svc) StatusCopy(ctx context.Context, req *game_pb.CopyRequest) (*game_pb.CopyResponse, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	log.Println(req)
	if s.role != models.BackupNode {
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_I_AM_NOT_BACKUP,
		}, nil
	}
	if s.gameCopy.GetStateVersion() <= req.GetStateVersion() {
		s.gameCopy = req
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_OK,
		}, nil
	} else {
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_NULL_ERROR,
		}, nil
	}
}

func (s *svc) TakeSlot(ctx context.Context, req *game_pb.TakeSlotRequest) (*game_pb.TakeSlotResponse, error) {
	log.Println(req)
	// only take requests when i am a primary server node
	if s.role == models.BackupNode {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_I_AM_ONLY_BACKUP,
		}, nil
	} else if s.role == models.PlayerNode {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_I_AM_NOT_A_SERVER,
		}, nil
	}

	// only master will deal with TakeSlot requests
	// 1. update locally
	// 2. serialise game status and sync with slave
	// 3. reply player

	// slave healthcheck
	if s.slave == nil {
		s.roleSetup()
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_SLAVE_INIT_IN_PROGRESS,
		}, nil
	}
	// dummy game state sync
	syncResp, syncErr := s.slave.StatusCopy(ctx, s.game.GetSerialisedGameStats())
	if syncErr != nil || syncResp.GetStatus() != game_pb.CopyResponse_OK {
		log.Printf("syncing with slave not successful: %s, %v", syncResp.GetStatus().String(), syncErr)
		// refresh slave
		s.roleSetup()
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_SLAVE_INIT_IN_PROGRESS,
		}, nil
	}

	row := int(req.GetMoveToCoordinate().GetRow())
	col := int(req.GetMoveToCoordinate().GetCol())

	_, err := s.game.PlacePlayer(req.GetId(), row, col)
	if err == models.InvalidCoordinates {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_INVALID_INPUT,
		}, nil
	} else if err == models.PlaceAlreadyTaken {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_SLOT_TAKEN,
		}, nil
	} else if err == models.SlaveIsDown {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_SLAVE_INIT_IN_PROGRESS,
		}, nil
	} else if err != nil {
		// unknown error occurred. e.g. io timeout. caller should retry
		return nil, err
	}

	// real sync. if failure happens, the next request will detect
	s.slave.StatusCopy(ctx, s.game.GetSerialisedGameStats())

	resp := &game_pb.TakeSlotResponse{
		Status: game_pb.TakeSlotResponse_OK,
	}

	gameStates := s.game.GetGameStates()
	for _, p := range gameStates {
		if p == nil {
			continue
		}
		resp.PlayerStates = append(resp.PlayerStates, p.ToPlayerProto())
	}

	return resp, nil
}

func (s *svc) Heartbeat(ctx context.Context, req *game_pb.HeartbeatRequest) (*game_pb.HeartbeatResponse, error) {
	s.rwLock.RLock()

	// if registry changes, there might be a role change as well
	log.Println(fmt.Sprintf("player %s received heartbeat from %s, registry version: %d", s.playerId, req.PlayerId, req.Registry.GetVersion()))
	if s.registry.GetVersion() < req.GetRegistry().GetVersion() {
		s.registry = req.GetRegistry()
		defer s.roleSetup()
	}

	defer s.rwLock.RUnlock()
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
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.role = s.deriveRole()
	switch s.role {
	case models.PrimaryNode:
		backupNode := GetBackupServer(s.registry)
		s.slave = ConnectToPlayer(backupNode)
	}
	return
}

func (s *svc) UpdateLocalRegistry(registry *game_pb.Registry) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.registry = registry
}

func (s *svc) GetLocalRegistry() *game_pb.Registry {
	//s.rwLock.RLock()
	//defer s.rwLock.Unlock()

	return s.registry
}

func NewGameSvc(playerId string, gridSize int, treasureAmount int, registry *game_pb.Registry) GameService {
	s := &svc{
		playerId: playerId,
		registry: registry,
		rwLock:   &sync.RWMutex{},
	}
	s.roleSetup()
	s.game = models.NewGame(gridSize, treasureAmount, s.role)
	return s
}
