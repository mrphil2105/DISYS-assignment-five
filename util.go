package main

import (
	"main/auction"
	"net"
	"strconv"

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

func (*Server) ConnectRingClient(port string) auction.RingServiceClient {
	log.Printf("Dialing ring node on port %s", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial ring node: %v", err)
	}

	return auction.NewRingServiceClient(conn)
}

func (*Frontend) ConnectServerClient(port string) auction.AuctionServiceClient {
	log.Printf("Dialing server on port %s", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	return auction.NewAuctionServiceClient(conn)
}

func ParsePort(portStr string) uint16 {
	parsed, err := strconv.ParseUint(portStr, 10, 16)

	if err != nil {
		log.Fatalf("Failed to parse '%s' as port: %v", portStr, err)
	}

	return uint16(parsed)
}
