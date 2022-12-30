package main

import (
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectClient(name string, port string) *grpc.ClientConn {
	log.Printf("Dialing %s on port %s", name, port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %s: %v", name, err)
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
