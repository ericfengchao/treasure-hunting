package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service/models"
)

type svc struct {
	role     models.Role
	playerId string
	registry *game_pb.Registry

	game     models.Gamer
	gameCopy *game_pb.CopyRequest
	slave    game_pb.GameServiceClient
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.game.GetGridView())
}

func (s *svc) StatusCopy(ctx context.Context, req *game_pb.CopyRequest) (*game_pb.CopyResponse, error) {
	log.Println(req)
	if s.role != models.BackupNode {
		return &game_pb.CopyResponse{
			Status: game_pb.CopyResponse_I_AM_NOT_BACKUP,
		}, nil
	}
	s.gameCopy = req
	return &game_pb.CopyResponse{
		Status: game_pb.CopyResponse_OK,
	}, nil
}

func (s *svc) MovePlayer(ctx context.Context, req *game_pb.MoveRequest) (*game_pb.MoveResponse, error) {
	log.Println(req)
	// only take requests when i am a primary server node
	if s.role == models.BackupNode {
		return &game_pb.MoveResponse{
			Status: game_pb.MoveResponse_I_AM_ONLY_BACKUP,
		}, nil
	} else if s.role == models.PlayerNode {
		return &game_pb.MoveResponse{
			Status: game_pb.TakeSlotResponse_I_AM_NOT_A_SERVER,
		}, nil
	}

	move := req.GetMove()
	Id := req.GetId()
	_, err := s.game.MovePlayer(Id, move)
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

	resp := &game_pb.MoveResponse{
		Status: game_pb.MoveResponse_OK,
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
	log.Println(fmt.Sprintf("player %s received heartbeat from %s, registry version: %d", s.playerId, req.PlayerId, req.Registry.GetVersion()))
	if s.registry.GetVersion() < req.GetRegistry().GetVersion() {
		s.registry = req.GetRegistry()
	}
	return &game_pb.HeartbeatResponse{}, nil
}

func (s *svc) UpdateLocalRegistry(registry *game_pb.Registry) {
	s.registry = registry
}

func (s *svc) GetLocalRegistry() *game_pb.Registry {
	return s.registry
}

func (s *svc) AttachSlave(slave game_pb.GameServiceClient) {
	s.slave = slave
}

func NewGameSvc(role models.Role, playerId string, gridSize int, treasureAmount int, registry *game_pb.Registry) GameService {
	return &svc{
		role:     role,
		playerId: playerId,
		registry: registry,
		game:     models.NewGame(gridSize, treasureAmount, role),
	}
}
