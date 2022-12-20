package main

import (
	"main/auction"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (*Server) ConnectBackupClient(port string) (auction.ConnectServiceClient, auction.AuctionServiceClient) {
	conn := ConnectClient("backup node", port)

	connectClient := auction.NewConnectServiceClient(conn)
	auctionClient := auction.NewAuctionServiceClient(conn)

	return connectClient, auctionClient
}

func (*Server) ConnectRingClient(port string) auction.RingServiceClient {
	conn := ConnectClient("ring node", port)

	return auction.NewRingServiceClient(conn)
}

func (*Frontend) ConnectServerClient(port string) auction.AuctionServiceClient {
	conn := ConnectClient("server", port)

	return auction.NewAuctionServiceClient(conn)
}

func ConnectClient(name string, port string) *grpc.ClientConn {
	log.Printf("Dialing %s on port %s", name, port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial frontend: %v", err)
	}

	return conn
}

func ParsePort(portStr string) uint16 {
	parsed, err := strconv.ParseUint(portStr, 10, 16)

	if err != nil {
		log.Fatalf("Failed to parse '%s' as port: %v", portStr, err)
	}

	return uint16(parsed)
}
