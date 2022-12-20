package main

import (
	"context"
	"main/auction"
	"main/ring"
)

// Called by gRPC
func (server *Server) SendElection(ctx context.Context, election *auction.Election) (*auction.Void, error) {
	ring.ReceiveElection(server, election.GetPid())

	return &auction.Void{}, nil
}

// Called by gRPC
func (server *Server) SendElected(ctx context.Context, elected *auction.Elected) (*auction.Void, error) {
	ring.ReceiveElected(server, elected.GetPid())

	return &auction.Void{}, nil
}

// Implement Ring
func (server *Server) GetPid() uint32 {
	return server.pid
}

// Implement Ring
func (server *Server) GetState() ring.State {
	return server.state
}

// Implement Ring
func (server *Server) SetState(state ring.State) {
	log.Printf("Setting Ring state to %s", state)
	server.state = state
}

// Implement Ring
func (server *Server) SetElected(pid uint32) {
	log.Printf("Setting elected to pid %d", pid)
	server.elected = pid
}

// Implement Ring
func (server *Server) Election(pid uint32) {
	if server.neighborClient == nil {
		log.Fatalf("Unable to send Election message with no neighbor connection")
	}

	log.Printf("Sending Election message with pid %d to port %s", pid, server.neighbor.port)

	_, err := server.neighborClient.SendElection(context.Background(), &auction.Election{Pid: pid})

	if err != nil {
		log.Fatalf("Failed to send Election message to neighbor: %v", err)
	}
}

// Implement Ring
func (server *Server) Elected(pid uint32) {
	if server.neighborClient == nil {
		log.Fatalf("Unable to send Election message with no neighbor connection")
	}

	log.Printf("Sending Elected message with pid %d to port %s", pid, server.neighbor.port)

	_, err := server.neighborClient.SendElected(context.Background(), &auction.Elected{Pid: pid})

	if err != nil {
		log.Fatalf("Failed to send Elected message to neighbor: %v", err)
	}
}
