package service

import (
	"errors"
	"fmt"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"github.com/ericfengchao/treasure-hunting/service/models"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
)

func ConnectToPlayer(player *game_pb.Player) game_pb.GameServiceClient {
	if player == nil {
		return nil
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", player.GetIp(), player.GetPort()),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("err connecting to player", player, err)
		return nil
	}
	return game_pb.NewGameServiceClient(conn)
}

func GetPrimaryServer(registry *game_pb.Registry) *game_pb.Player {
	if len(registry.GetPlayerList()) > 0 {
		return registry.GetPlayerList()[0]
	}
	return nil
}

func GetBackupServer(registry *game_pb.Registry) *game_pb.Player {
	if len(registry.GetPlayerList()) > 1 {
		return registry.GetPlayerList()[1]
	}
	return nil
}

var InvalidMoveInput = errors.New("invalid move input")

func ParseDirection(input string) (models.Movement, error) {
	i, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return models.Stay, err
	}
	switch i {
	case 0, 1, 2, 3, 4, 9:
		return models.Movement(i), nil
	default:
		return models.Stay, InvalidMoveInput
	}
}
