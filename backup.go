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

func NewBackup(pid uint32, port string) *Backup {
	return &Backup{
		pid:  pid,
		port: port,
	}
}

func (backup *Backup) SetConnection(connectClient auction.ConnectServiceClient,
	auctionClient auction.AuctionServiceClient) {

	backup.connectClient = connectClient
	backup.auctionClient = auctionClient
}

func (backup *Backup) ClearConnection() {
	backup.connectClient = nil
	backup.auctionClient = nil
}

// Called by gRPC
func (*Server) FinishConnect(ctx context.Context, void *auction.Void) (*auction.BackupDetails, error) {
	return &auction.BackupDetails{Pid: uint32(os.Getpid())}, nil
}

// Called by gRPC
func (server *Server) AddBackup(ctx context.Context, backupJoin *auction.BackupJoin) (*auction.Void, error) {
	pid := backupJoin.GetPid()
	port := backupJoin.GetPort()
	server.backups[pid] = NewBackup(pid, port)

	log.Printf("Received information about a backup (pid %d, port %s)", pid, port)

	return &auction.Void{}, nil
}

// Called by gRPC
func (server *Server) RemoveBackup(ctx context.Context, backupLeave *auction.BackupLeave) (*auction.Void, error) {
	pid := backupLeave.GetPid()
	backup := server.backups[pid]
	delete(server.backups, pid)

	log.Printf("Received information that a backup (pid %d, port %s) was removed", pid, backup.port)

	return &auction.Void{}, nil
}
