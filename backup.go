package main

import (
	"context"
	"main/auction"
	"os"
)

type Backup struct {
	pid  uint32
	port string
}

type BackupConn struct {
	backup        *Backup
	connectClient auction.ConnectServiceClient
	auctionClient auction.AuctionServiceClient
}

func NewBackup(pid uint32, port string) *Backup {
	return &Backup{
		pid:  pid,
		port: port,
	}
}

func NewBackupConn(backup *Backup, connectClient auction.ConnectServiceClient,
	auctionClient auction.AuctionServiceClient) *BackupConn {

	return &BackupConn{
		backup:        backup,
		connectClient: connectClient,
		auctionClient: auctionClient,
	}
}

// Called by gRPC
func (*Server) FinishConnect(ctx context.Context, void *auction.Void) (*auction.BackupDetails, error) {
	log.Printf("Sending backup details (pid %d) to main node", os.Getpid())

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

// Called by gRPC
func (server *Server) AuctionStarted(ctx context.Context, void *auction.Void) (*auction.Void, error) {
	if server.pid == server.elected {
		log.Fatalf("Received auction started message as main node")
	}

	server.bids = make(map[uint32]*Bid)
	server.auctionDone = false

	return &auction.Void{}, nil
}
