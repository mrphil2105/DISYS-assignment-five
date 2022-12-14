package main

import (
	"bufio"
	"context"
	"fmt"
	"main/auction"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	pid            uint32
	frontendClient auction.FrontendServiceClient
}

func NewClient() *Client {
	return &Client{
		pid: uint32(os.Getpid()),
	}
}

func client() {
	conn := ConnectClient("frontend", frontendPort)
	client := NewClient()
	client.frontendClient = auction.NewFrontendServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")

		switch input[0] {
		case "bet":
			SendBid(client, input)
		case "result":
			GetResult(client)
		default:
			fmt.Printf("Unknown command '%s'\n", input[0])
		}
	}
}

func SendBid(client *Client, input []string) {
	amount, err := strconv.ParseUint(input[1], 10, 32)

	if err != nil {
		log.Printf("Unable to parse '%s' as amount: %v", input[1], err)
		return
	}

	bid := &auction.Bid{
		Pid:    client.pid,
		Amount: uint32(amount),
	}
	ack, err := client.frontendClient.SendBid(context.Background(), bid)

	if err != nil {
		log.Fatalf("Unable to place bid: %v", err)
	}

	if ack.GetSuccess() {
		log.Printf("Bid of '%d' successfully placed", amount)
	} else {
		log.Printf("Bid of '%d' was not higher than the highest bid", amount)
	}
}

func GetResult(client *Client) {
	outcome, err := client.frontendClient.GetResult(context.Background(), &auction.Void{})

	if err != nil {
		log.Fatalf("Unable to get auction outcome: %v", err)
	}

	log.Printf("Highest bid: %d, winner pid: %d", outcome.GetHighestBid(), outcome.GetWinner())
	log.Printf("Bidders:")

	for _, bidder := range outcome.GetBidders() {
		log.Printf("  %d", bidder)
	}
}
