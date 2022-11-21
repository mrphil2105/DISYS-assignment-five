package main

import as "main/auctionSystem"

type Peer struct {
	pid           uint32
	connectClient as.ConnectServiceClient
	auctionClient as.AuctionServiceClient
}

func NewPeer(pid uint32, connectClient as.ConnectServiceClient, auctionClient as.AuctionServiceClient) *Peer {
	return &Peer{
		pid:           pid,
		connectClient: connectClient,
		auctionClient: auctionClient,
	}
}
