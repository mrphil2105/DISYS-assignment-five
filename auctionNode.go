package main

import (
	as "main/auctionSystem"
	"os"
)

type AuctionNode struct {
	as.UnimplementedConnectServiceServer
	as.UnimplementedAuctionServiceServer
	as.UnimplementedBullyServiceServer
	pid   uint32
	peers map[uint32]*Peer
}

func NewAuctionNode() *AuctionNode {
	return &AuctionNode{
		pid:   uint32(os.Getpid()),
		peers: make(map[uint32]*Peer),
	}
}

func (node *AuctionNode) GetPid() uint32 {
	return node.pid
}
