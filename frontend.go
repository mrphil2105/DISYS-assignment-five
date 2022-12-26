package main

import (
	"context"
	"google.golang.org/grpc"
	"main/auction"
	"net"
)

type Frontend struct {
	auction.UnimplementedFrontendServiceServer
	auctionClient auction.AuctionServiceClient
}

func NewFrontend() *Frontend {
	return &Frontend{}
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
	if frontend.auctionClient == nil {
		log.Fatalf("Unable to send bid when there is no main server set")
	}

	log.Printf("Sending bid (pid %d, amount %d) to main server", bid.GetPid(), bid.GetAmount())
	ack, err := frontend.auctionClient.SendBid(ctx, bid)

	if err != nil {
		log.Fatalf("Unable to send bid to main server: %v", err)
	}

	return ack, nil
}

// Called by gRPC
func (frontend *Frontend) GetResult(ctx context.Context, void *auction.Void) (*auction.Outcome, error) {
	if frontend.auctionClient == nil {
		log.Fatalf("Unable to get result when there is no main server set")
	}

	log.Printf("Getting result from main server")
	outcome, err := frontend.auctionClient.GetResult(ctx, void)

	if err != nil {
		log.Fatalf("Unable to get result from main server: %v", err)
	}

	return outcome, nil
}

// Called by gRPC
func (frontend *Frontend) SetPrimaryNode(ctx context.Context, primaryNode *auction.PrimaryNode) (*auction.Void, error) {
	pid := primaryNode.GetPid()
	port := primaryNode.GetPort()
	log.Printf("Setting main server to pid %d, port %s", pid, port)

	conn := ConnectClient("server", port)
	frontend.auctionClient = auction.NewAuctionServiceClient(conn)

	return &auction.Void{}, nil
}
