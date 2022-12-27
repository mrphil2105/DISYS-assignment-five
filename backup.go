package main

import (
	"context"
	"main/auction"
	"os"
)

type Backup struct {
	pid           uint32
	port          string
	connectClient auction.ConnectServiceClient
	auctionClient auction.AuctionServiceClient
}

func NewBackup(pid uint32, port string, connectClient auction.ConnectServiceClient,
	auctionClient auction.AuctionServiceClient) *Backup {

	return &Backup{
		pid:           pid,
		port:          port,
		connectClient: connectClient,
		auctionClient: auctionClient,
	}
}

// Called by gRPC
func (server *Server) FinishConnect(ctx context.Context, void *auction.Void) (*auction.BackupDetails, error) {
	log.Printf("Sending backup details (pid %d) to main node", os.Getpid())

	// TODO: Remove the main server from backups (in case the main server was a backup)

	server.StartPingTimer()

	return &auction.BackupDetails{Pid: uint32(os.Getpid())}, nil
}

// Called by gRPC
func (server *Server) AddBackup(ctx context.Context, backupJoin *auction.BackupJoin) (*auction.Void, error) {
	pid := backupJoin.GetPid()
	port := backupJoin.GetPort()

	backup := NewBackup(pid, port, nil, nil)
	server.SetBackup(backup)

	log.Printf("Received information about a backup (pid %d, port %s)", pid, port)

	return &auction.Void{}, nil
}

// Called by gRPC
func (server *Server) RemoveBackup(ctx context.Context, backupLeave *auction.BackupLeave) (*auction.Void, error) {
	pid := backupLeave.GetPid()
	backup := server.GetBackup(pid)
	server.DeleteBackup(pid)

	log.Printf("Received information that a backup (pid %d, port %s) was removed", pid, backup.port)

	return &auction.Void{}, nil
}

// Called by gRPC
func (server *Server) AuctionStarted(ctx context.Context, void *auction.Void) (*auction.Void, error) {
	if server.isMain {
		log.Fatalf("Received auction started message as main node")
	}

	server.bids = make(map[uint32]*Bid)
	server.auctionDone = false

	return &auction.Void{}, nil
}

// Called by gRPC
func (server *Server) Ping(ctx context.Context, void *auction.Void) (*auction.Void, error) {
	if server.isMain {
		log.Fatalf("Received ping message as main node")
	}

	server.pingChan <- true

	return &auction.Void{}, nil
}
