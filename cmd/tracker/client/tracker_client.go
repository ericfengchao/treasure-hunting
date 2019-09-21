package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"

	tracker_pb "github.com/ericfengchao/treasure-hunting/protos/tracker"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Wrong Args Number")
	}

	host := os.Args[1]     // tracker's ip address
	port := os.Args[2]     // tracker's port
	playerId := os.Args[3] // player's id

	address := host + ":" + port // concat address

	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to server", err)
	}
	defer conn.Close()
	client := tracker_pb.NewTrackerServiceClient(conn)
	resp, err := client.Register(context.Background(), &tracker_pb.RegisterRequest{
		PlayerId: playerId,
	})
	//resp2, _ := client.ReportMissing(context.Background(), &tracker_pb.Missing{
	//	PlayerId: playerId,
	//})
	jpb := jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "  ",
	}
	o, _ := jpb.MarshalToString(resp)
	fmt.Printf("Resp:\n%s\n", o)
	// log.Println(resp2)
}
