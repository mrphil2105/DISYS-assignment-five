package main

import (
	"context"
	"main/auction"
)

// Called by gRPC
func (server *Server) SendBid(ctx context.Context, bidMsg *auction.Bid) (*auction.BidAck, error) {
	pid := bidMsg.GetPid()
	amount := bidMsg.GetAmount()

	for _, bid := range server.bids {
		if bid.amount >= amount {
			return &auction.BidAck{Success: false}, nil
		}
	}

	server.bids[pid] = NewBid(pid, amount)
	log.Printf("Received bid from client (pid %d) with amount: %d", pid, amount)

	// Send to backups if we are the main node.
	if server.elected == server.pid {
		for _, backupConn := range server.backupConns {
			log.Printf("Sending bid to backup (pid %d, port %s)", backupConn.backup.pid, backupConn.backup.port)
			_, err := backupConn.auctionClient.SendBid(context.Background(), bidMsg)

			if err != nil {
				log.Printf("Failed to send bid to backup (pid %d, port %s)", backupConn.backup.pid,
					backupConn.backup.port)
			}
		}
	}

	return &auction.BidAck{Success: true}, nil
}

// Called by gRPC
func (server *Server) GetResult(ctx context.Context, void *auction.Void) (*auction.Outcome, error) {
	pids := make([]uint32, 0, len(server.bids))
	var winner, highestBid uint32

	for pid, bid := range server.bids {
		pids = append(pids, pid)

		if bid.amount > highestBid {
			highestBid = bid.amount

			if server.auctionDone {
				winner = pid
			}
		}
	}

	log.Printf("Sending auction outcome (bidders %v, winner %d, highest bid %d)", pids, winner, highestBid)

	return &auction.Outcome{
		Bidders:    pids,
		Winner:     winner,
		HighestBid: highestBid,
	}, nil
}
