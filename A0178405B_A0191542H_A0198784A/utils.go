package treasure_hunting

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
)

func ConnectToPlayer(player *Player) (*grpc.ClientConn, GameServiceClient) {
	if player == nil {
		return nil, nil
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", player.GetIp(), player.GetPort()),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("err connecting to player", player, err)
		return nil, nil
	}
	return conn, NewGameServiceClient(conn)
}

func GetPrimaryServer(registry *Registry) *Player {
	if len(registry.GetPlayerList()) > 0 {
		return registry.GetPlayerList()[0]
	}
	return nil
}

func GetBackupServer(registry *Registry) *Player {
	if len(registry.GetPlayerList()) > 1 {
		return registry.GetPlayerList()[1]
	} else if len(registry.GetPlayerList()) == 1 {
		return registry.GetPlayerList()[0]
	}
	return nil
}

var InvalidMoveInput = errors.New("invalid move input")

func ParseDirection(input string) (Movement, error) {
	i, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return Stay, err
	}
	switch i {
	case 0, 1, 2, 3, 4, 9:
		return Movement(i), nil
	default:
		return Stay, InvalidMoveInput
	}
}
