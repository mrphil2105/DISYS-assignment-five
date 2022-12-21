package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"main/auction"
	"main/ring"
	"net"
	"os"
	"strconv"
	"strings"
)

type Bid struct {
	pid    uint32
	amount uint32
}

type Server struct {
	auction.UnimplementedConnectServiceServer
	auction.UnimplementedAuctionServiceServer
	auction.UnimplementedRingServiceServer
	pid            uint32
	bids           map[uint32]*Bid
	auctionDone    bool
	state          ring.State
	elected        uint32
	neighbor       *Backup
	neighborClient auction.RingServiceClient
	backups        map[uint32]*Backup
	backupConns    map[uint32]*BackupConn
}

func NewServer() *Server {
	return &Server{
		pid:         uint32(os.Getpid()),
		bids:        make(map[uint32]*Bid),
		backups:     make(map[uint32]*Backup),
		backupConns: make(map[uint32]*BackupConn),
	}
}

func NewBid(pid uint32, amount uint32) *Bid {
	return &Bid{
		pid:    pid,
		amount: amount,
	}
}

func server() {
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", port))

	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", port, err)
	}

	defer listener.Close()

	server := NewServer()

	go func() {
		grpcServer := grpc.NewServer()
		auction.RegisterConnectServiceServer(grpcServer, server)
		auction.RegisterAuctionServiceServer(grpcServer, server)
		log.Printf("Created gRPC server on port %s", port)

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Stopped serving due to error: %v", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		switch input[0] {
		case "connect":
			ConnectToBackup(server, input)
		case "kill":
			os.Exit(0)
		default:
			fmt.Printf("Unknown command '%s'\n", input[0])
		}
	}
}

func ConnectToBackup(server *Server, input []string) {
	if len(server.backups) > 0 && server.elected != server.pid {
		log.Printf("Adding a backup connection to a backup is not valid")
		return
	}

	port := strconv.Itoa(int(countingPort + ParsePort(input[1])))
	conn := ConnectClient("backup node", port)

	connectClient := auction.NewConnectServiceClient(conn)
	auctionClient := auction.NewAuctionServiceClient(conn)

	details, err := connectClient.FinishConnect(context.Background(), &auction.Void{})

	if err != nil {
		log.Fatalf("Failed to get backup details from %s: %v", port, err)
	}

	log.Printf("Informing backups about new backup (pid %d, port %s)", details.GetPid(), port)

	InformExistingBackups(server, details.GetPid(), port)
	InformNewBackup(server, connectClient)

	backup := NewBackup(details.GetPid(), port)
	backupConn := NewBackupConn(backup, connectClient, auctionClient)

	server.SetAsMain()
	server.backups[details.GetPid()] = backup
	server.backupConns[details.GetPid()] = backupConn
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

func (server *Server) SetAsMain() {
	if server.elected != server.pid {
		log.Printf("Setting server as main node")
		server.elected = server.pid

		conn := ConnectClient("frontend", frontendPort)
		defer conn.Close()

		frontendClient := auction.NewFrontendServiceClient(conn)

		elected := &auction.Elected{
			Pid:  server.pid,
			Port: port,
		}
		_, err := frontendClient.SetPrimaryNode(context.Background(), elected)

		if err != nil {
			log.Fatalf("Unable to inform frontend about main node")
		}
	}
}

func (server *Server) GetRingNeighbor() *Backup {
	backupsByPort := make(map[uint16]*Backup)
	var lowestPort uint16 = 0xFFFF
	var highestPort uint16

	for _, backup := range server.backups {
		backupPort := ParsePort(backup.port)
		backupsByPort[backupPort] = backup

		if backupPort < lowestPort {
			lowestPort = backupPort
		}

		if backupPort > highestPort {
			highestPort = backupPort
		}
	}

	port := ParsePort(port) + 1

	if port > highestPort {
		port = lowestPort
	}

	return backupsByPort[port]
}

func (server *Server) ConnectToRingNeighbor() {
	conn := ConnectClient("ring node", port)
	server.neighbor = server.GetRingNeighbor()
	server.neighborClient = auction.NewRingServiceClient(conn)
}
