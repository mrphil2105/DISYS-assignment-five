package main

import (
	"main/auction"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (*Server) ConnectBackupClient(port string) (auction.ConnectServiceClient, auction.AuctionServiceClient) {
	log.Printf("Dialing node on port %s", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial node: %v", err)
	}

	connectClient := auction.NewConnectServiceClient(conn)
	auctionClient := auction.NewAuctionServiceClient(conn)

	return connectClient, auctionClient
}
