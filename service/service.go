package service

import (
	"context"
	"fmt"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service/models"
	"log"
	"net/http"
	"time"
)

type svc struct {
	role     Role
	playerId string
	registry *game_pb.Registry

	// heartbeat neighbour
	prevNeighbour game_pb.GameServiceClient
	nextNeighbour game_pb.GameServiceClient

	game  models.Gamer
	slave game_pb.GameServiceClient
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.game.GetGridView())
}

func (s *svc) TakeSlot(ctx context.Context, req *game_pb.TakeSlotRequest) (*game_pb.TakeSlotResponse, error) {
	log.Println(req)
	// only take requests when i am a primary server node
	if s.role == BackupNode {
		return &game_pb.TakeSlotResponse{
			Status: game_pb.TakeSlotResponse_I_AM_ONLY_BACKUP,
		}, nil
	} else if s.role == PlayerNode {
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
	if s.registry.GetVersion() < req.GetRegistry().GetVersion() {
		s.registry = req.GetRegistry()
		// update heartbeating neighbours
		// if node's role is primary, contact new backup and sync
		// if node's role is secondary, self update to primary and contact new secondary and sync
		// if node's role is player, self update to backup and wait for primary to sync
	}
	return &game_pb.HeartbeatResponse{}, nil
}

func (s *svc) AttachSlave(slave game_pb.GameServiceClient) {
	s.slave = slave
}

func (s *svc) HeartbeatInBackground(ctx context.Context) {
	t := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-t.C:
			ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*200)
			func() {
				defer cancelFunc()
				req := &game_pb.HeartbeatRequest{
					PlayerId: s.playerId,
					Registry: s.registry,
				}
				if s.nextNeighbour != nil {
					_, err := s.nextNeighbour.Heartbeat(ctx, req)
					if err != nil {
						// report next neighbour down
					}
				}
				if s.prevNeighbour != nil {
					_, err := s.prevNeighbour.Heartbeat(ctx, req)
					if err != nil {
						// report prev neighbour down
					}
				}
			}()
		}
	}
}

func NewGameSvc(role Role, playerId string, gridSize int, treasureAmount int, registry *game_pb.Registry) GameService {
	return &svc{
		role:     role,
		playerId: playerId,
		registry: registry,
		game:     models.NewGame(gridSize, treasureAmount),
	}
}
