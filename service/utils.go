package service

import (
	"fmt"
	game_pb "github.com/ericfengchao/treasure-hunting/protos"
	"google.golang.org/grpc"
	"log"
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
