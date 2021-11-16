package main

import (
	"context"
	"log"
	"time"

	proto "github.com/JoakimGD/MiniProject3/auctionhouse"
	"google.golang.org/grpc"
)

var (
	id int64 = 3
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()
	c := proto.NewAuctionHouseClient(conn)

	for {
		r1, err := c.RequestToken(context.Background(), &proto.Node{Id: id})
		if err != nil {
			log.Fatalf("RequestToken failed: %s", err)
		}
		log.Printf("Node: %d entered CS", r1.GetFrom())
		time.Sleep(6 * time.Second)

		r2, err := c.ReturnToken(context.Background(), &proto.Token{From: id})
		if err != nil {
			log.Fatalf("ReturnToken failed: %s", err)
		}
		log.Printf("Node: %d left CS", r2.GetFrom())
		time.Sleep(2 * time.Second) // Added 2 sec to better see switch between nodes in CS
	}

}
