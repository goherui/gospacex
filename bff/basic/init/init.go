package init

import (
	"flag"
	"gospaacex/bff/basic/config"
	__ "gospaacex/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	InitGRpc()
}
func InitGRpc() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.PosClient = __.NewStreamGreeterClient(conn)
}
