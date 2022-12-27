package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"main/auction"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Bid struct {
	pid    uint32
	amount uint32
}

type Server struct {
	auction.UnimplementedConnectServiceServer
	auction.UnimplementedAuctionServiceServer
	pid            uint32
	bids           map[uint32]*Bid
	auctionDone    bool
	isMain         bool
	pingChan       chan interface{}
	pingTimer      bool
	pingTimerMutex sync.Mutex
	backups        map[uint32]*Backup
	backupsMutex   sync.Mutex
}

func NewServer() *Server {
	return &Server{
		pid:      uint32(os.Getpid()),
		bids:     make(map[uint32]*Bid),
		pingChan: make(chan interface{}, 1),
		backups:  make(map[uint32]*Backup),
	}
}

func NewBid(pid uint32, amount uint32) *Bid {
	return &Bid{
		pid:    pid,
		amount: amount,
	}
}

func (server *Server) GetBackups() map[uint32]*Backup {
	backups := make(map[uint32]*Backup)

	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	for pid, backup := range server.backups {
		backups[pid] = backup
	}

	return backups
}

func (server *Server) GetBackup(pid uint32) *Backup {
	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	return server.backups[pid]
}

func (server *Server) GetBackupCount() int {
	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	return len(server.backups)
}

func (server *Server) SetBackup(backup *Backup) {
	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	server.backups[backup.pid] = backup
}

func (server *Server) DeleteBackup(pid uint32) {
	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	delete(server.backups, pid)
}

func (server *Server) ClearBackups() {
	server.backupsMutex.Lock()
	defer server.backupsMutex.Unlock()

	server.backups = make(map[uint32]*Backup)
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
			if server.GetBackupCount() > 0 && !server.isMain {
				log.Printf("Adding a backup connection to a backup is not valid")
				return
			}

			port := strconv.Itoa(int(countingPort + ParsePort(input[1])))
			ConnectToBackup(server, port)
		case "kill":
			os.Exit(0)
		default:
			fmt.Printf("Unknown command '%s'\n", input[0])
		}
	}
}

func ConnectToBackup(server *Server, port string) {
	conn := ConnectClient("backup node", port)

	connectClient := auction.NewConnectServiceClient(conn)
	auctionClient := auction.NewAuctionServiceClient(conn)

	details, err := connectClient.FinishConnect(context.Background(), &auction.Void{})

	if err != nil {
		log.Fatalf("Failed to get backup details from %s: %v", port, err)
	}

	log.Printf("Informing backups about new backup (pid %d, port %s)", details.GetPid(), port)

	InformBackupsAboutJoin(server, details.GetPid(), port)
	InformNewBackup(server, connectClient)

	backup := NewBackup(details.GetPid(), port, connectClient, auctionClient)
	server.SetAsMain()
	server.SetBackup(backup)
}

func InformBackupsAboutJoin(server *Server, pid uint32, port string) {
	// Inform backups about the new backup.
	for _, backup := range server.GetBackups() {
		_, err := backup.connectClient.AddBackup(context.Background(), &auction.BackupJoin{
			Pid:  pid,
			Port: port,
		})

		if err != nil {
			log.Fatalf("Failed to inform backup (pid %d, port %s) about new backup", backup.pid, backup.port)
		}
	}
}

func InformBackupsAboutLeave(server *Server, pid uint32) {
	// Inform backups about the backup removal.
	for _, backup := range server.GetBackups() {
		_, err := backup.connectClient.RemoveBackup(context.Background(), &auction.BackupLeave{Pid: pid})

		if err != nil {
			log.Fatalf("Failed to inform backup (pid %d, port %s) about removed backup", backup.pid, backup.port)
		}
	}
}

func InformNewBackup(server *Server, connectClient auction.ConnectServiceClient) {
	// Inform the new backup about existing backups.
	for pid, backup := range server.GetBackups() {
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
	if !server.isMain {
		log.Printf("Setting server as main node")
		server.isMain = true

		conn := ConnectClient("frontend", frontendPort)
		defer conn.Close()

		frontendClient := auction.NewFrontendServiceClient(conn)

		primaryNode := &auction.PrimaryNode{
			Pid:  server.pid,
			Port: port,
		}
		_, err := frontendClient.SetPrimaryNode(context.Background(), primaryNode)

		if err != nil {
			log.Fatalf("Unable to inform frontend about main node")
		}

		server.StartHeartbeat()
	}
}

func (server *Server) StartHeartbeat() {
	if !server.isMain {
		log.Fatalf("Unable to start heartbeat on a backup server")
	}

	go func() {
		for {
			time.Sleep(time.Second)

			func() {
				server.backupsMutex.Lock()
				defer server.backupsMutex.Unlock()

				for _, backup := range server.backups {
					_, err := backup.auctionClient.Ping(context.Background(), &auction.Void{})

					if err != nil {
						pid := backup.pid

						log.Printf("Unable to send ping to backup (pid %d, port %s): %v", pid, backup.port, err)
						log.Printf("Removing backup (pid %d, port %s)", pid, backup.port)

						delete(server.backups, pid)
						go InformBackupsAboutLeave(server, pid)
					}
				}
			}()
		}
	}()
}

func (server *Server) StartPingTimer() {
	if server.isMain {
		log.Fatalf("Unable to start ping timer on a main server")
	}

	shouldContinue := func() bool {
		server.pingTimerMutex.Lock()
		defer server.pingTimerMutex.Unlock()

		if server.pingTimer {
			return false
		}

		server.pingTimer = true
		return true
	}()

	if !shouldContinue {
		return
	}

	go func() {
		for {
			shouldStop := func() bool {
				server.pingTimerMutex.Lock()
				defer server.pingTimerMutex.Unlock()

				return !server.pingTimer
			}()

			if shouldStop {
				return
			}

			afterChan := time.After(3 * time.Second)

			select {
			case <-server.pingChan:
				break
			case <-afterChan:
				log.Printf("No Ping message received in the last 3 seconds")
				server.StopPingTimer()

				highestPid := server.pid

				for _, backup := range server.GetBackups() {
					if backup.pid > highestPid {
						highestPid = backup.pid
					}
				}

				if highestPid == server.pid {
					server.TakeOverAsMain()
				}
			}
		}
	}()
}

func (server *Server) StopPingTimer() {
	server.pingTimerMutex.Lock()
	defer server.pingTimerMutex.Unlock()

	server.pingTimer = false
}

func (server *Server) TakeOverAsMain() {
	log.Printf("Taking over as main node")
	backups := server.GetBackups()
	server.ClearBackups()

	for _, backup := range backups {
		ConnectToBackup(server, backup.port)
	}
}
