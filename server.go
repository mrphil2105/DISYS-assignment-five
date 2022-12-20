package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"main/auction"
	"net"
	"os"
	"strings"
)

type Bid struct {
	pid    uint32
	amount uint32
}

type Server struct {
	auction.UnimplementedConnectServiceServer
	pid         uint32
	bids        map[uint32]*Bid
	auctionDone bool
	elected     uint32
	backups     map[uint32]*Backup
	backupConns map[uint32]*BackupConnection
}

func NewServer() *Server {
	return &Server{
		pid:         uint32(os.Getpid()),
		bids:        make(map[uint32]*Bid),
		backups:     make(map[uint32]*Backup),
		backupConns: make(map[uint32]*BackupConnection),
	}
}

func NewBid(pid uint32, amount uint32) *Bid {
	return &Bid{
		pid:    pid,
		amount: amount,
	}
}

func server() {
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", *port))

	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", *port, err)
	}

	defer listener.Close()

	server := NewServer()

	go func() {
		grpcServer := grpc.NewServer()
		auction.RegisterConnectServiceServer(grpcServer, server)
		log.Printf("Created gRPC server on port %s", *port)

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Stopped serving due to error: %v", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		switch input[0] {
		case "connect":
			clientPort := input[1]
			connectClient, auctionClient := server.ConnectBackupClient(clientPort)
			details, err := connectClient.FinishConnect(context.Background(), &auction.Void{})

			if err != nil {
				log.Fatalf("Failed to get backup details from %s: %v", clientPort, err)
			}

			log.Printf("Informing backups about new backup (pid %d, port %s)", details.GetPid(), clientPort)

			InformExistingBackups(server, details.GetPid(), clientPort)
			InformNewBackup(server, connectClient)

			backup := NewBackup(details.GetPid(), clientPort)
			backupConn := NewBackupConnection(backup, connectClient, auctionClient)

			server.backups[details.GetPid()] = backup
			server.backupConns[details.GetPid()] = backupConn
		case "kill":
			os.Exit(0)
		default:
			fmt.Printf("Unknown command '%s'\n", input[0])
		}
	}
}

func InformExistingBackups(server *Server, newPid uint32, newPort string) {
	// Inform existing backups about the new backup.
	for pid, backupConn := range server.backupConns {
		_, err := backupConn.connectClient.AddBackup(context.Background(), &auction.BackupJoin{
			Pid:  newPid,
			Port: newPort,
		})

		if err != nil {
			log.Fatalf("Failed to inform backup (pid %d, port %s) about new backup", pid, backupConn.backup.port)
		}
	}
}

func InformNewBackup(server *Server, connectClient auction.ConnectServiceClient) {
	// Inform the new backup about existing backups.
	for pid, backup := range server.backups {
		_, err := connectClient.AddBackup(context.Background(), &auction.BackupJoin{
			Pid:  pid,
			Port: backup.port,
		})

		if err != nil {
			log.Fatalf("Failed to inform new backup about backup (pid %d, port %s)", pid, backup.port)
		}
	}
}
