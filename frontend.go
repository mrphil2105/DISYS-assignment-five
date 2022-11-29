package main

import (
	"context"
	"google.golang.org/grpc"
	"main/auction"
	"net"
)

type ServerConnection struct {
	pid    uint32
	port   string
	client auction.AuctionServiceClient
}

type Frontend struct {
	auction.UnimplementedFrontendServiceServer
	server *ServerConnection
}

func NewFrontend() *Frontend {
	return &Frontend{}
}

func NewServerConnection(pid uint32, port string, client auction.AuctionServiceClient) *ServerConnection {
	return &ServerConnection{
		pid:    pid,
		port:   port,
		client: client,
	}
}

func frontend() {
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", frontendPort))

	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", frontendPort, err)
	}

	defer listener.Close()

	frontend := NewFrontend()

	grpcServer := grpc.NewServer()
	auction.RegisterFrontendServiceServer(grpcServer, frontend)
	log.Printf("Created gRPC server on port %s", frontendPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Stopped serving due to error: %v", err)
	}
}

// Called by gRPC
func (frontend *Frontend) SendBid(ctx context.Context, bid *auction.Bid) (*auction.BidAck, error) {
	log.Printf("Sending bid (pid %d, amount %d) to main server", bid.GetPid(), bid.GetAmount())

	server := frontend.GetServer()
	ack, err := server.client.SendBid(ctx, bid)

	if err != nil {
		log.Fatalf("Unable to send bid to main server: %v", err)
	}

	return ack, nil
}

// Called by gRPC
func (frontend *Frontend) GetResult(ctx context.Context, void *auction.Void) (*auction.Outcome, error) {
	log.Printf("Getting result from main server")

	server := frontend.GetServer()
	outcome, err := server.client.GetResult(ctx, void)

	if err != nil {
		log.Fatalf("Unable to get result from main server: %v", err)
	}

	return outcome, nil
}

// Called by gRPC
func (frontend *Frontend) SetPrimaryNode(ctx context.Context, elected *auction.Elected) (*auction.Void, error) {
	pid := elected.GetPid()
	port := elected.GetPort()
	log.Printf("Setting main server to pid %d, port %s", pid, port)

	client := frontend.ConnectServerClient(port)
	frontend.server = NewServerConnection(pid, port, client)

	return &auction.Void{}, nil
}

func (frontend *Frontend) GetServer() *ServerConnection {
	if server := frontend.server; server != nil {
		return server
	}

	log.Fatalf("There is no main server set")
	return nil
}
