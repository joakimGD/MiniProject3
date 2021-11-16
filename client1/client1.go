package main

import (
	"context"
	"log"
	"time"

	proto "github.com/JoakimGD/MiniProject3/auctionhouse"
	"google.golang.org/grpc"
)

var (
	id int64 = 1
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()
	c := proto.NewAuctionHouseClient(conn)

	for {
		r1, err := c.Bid(context.Background(), 5) //needs to be amount value
		if err != nil {
			log.Fatalf("Bid failed: %s", err)
		}
		log.Printf("Node: %d entered CS", r1.GetSuccess()) //r1.GetSuccess() er ren og skær test pt
		time.Sleep(6 * time.Second)

		r2, err := c.Result(context.Background(), &proto.Outcome{})
		if err != nil {
			log.Fatalf("Request failed: %s", err)
		}
		log.Printf("Node: %d left CS", r2.GetHighestBid())
	}

}
