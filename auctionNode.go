package main

import (
	"context"
	"fmt"
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

// JoinNetwork ConnectService peer join implementation
func (node *AuctionNode) JoinNetwork(ctx context.Context, peerJoin *as.PeerJoin) (*as.ConnectedTo, error) {
	_, hasPeer := node.peers[peerJoin.GetPid()]

	if !hasPeer {
		connectClient, auctionClient, err := node.ConnectClient(peerJoin.GetPort())

		if err != nil {
			return nil, err
		}

		node.AddPeer(NewPeer(peerJoin.GetPid(), connectClient, auctionClient))
		log.Printf("Connected to peer %d", peerJoin.GetPid())
		fmt.Printf("New peer connected (total: %d)\n", len(node.peers))

		// Inform the rest of the network about the new peer.
		for pid, peer := range node.peers {
			// Check whether the current peer is the new peer.
			if pid != peerJoin.GetPid() {
				// Tell the other peer about the new peer.
				_, err = peer.connectClient.JoinNetwork(ctx, peerJoin)
			} else {
				// Respond to the new peer telling them about us.
				_, err = peer.connectClient.JoinNetwork(ctx, &as.PeerJoin{
					Pid:  node.GetPid(),
					Port: *port,
				})
			}

			if err != nil {
				log.Fatalf("Failed to propagate join: %v", err)
			}
		}
	}

	return &as.ConnectedTo{Pid: node.GetPid()}, nil
}

func (node *AuctionNode) AddPeer(peer *Peer) {
	node.peers[peer.pid] = peer
}

func (node *AuctionNode) RemovePeer(peer *Peer) {
	delete(node.peers, peer.pid)
}

func (node *AuctionNode) GetPid() uint32 {
	return node.pid
}
