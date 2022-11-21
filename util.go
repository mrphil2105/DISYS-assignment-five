package main

import (
	as "main/auctionSystem"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (node *AuctionNode) ConnectClient(port string) (as.ConnectServiceClient, as.AuctionServiceClient, error) {
	log.Printf("Dialing node on port %s...", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial node: %v", err)

		return nil, nil, err
	}

	connectClient := as.NewConnectServiceClient(conn)
	auctionClient := as.NewAuctionServiceClient(conn)

	return connectClient, auctionClient, nil
}
