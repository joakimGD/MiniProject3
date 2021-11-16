package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"

	proto "github.com/JoakimGD/MiniProject3/auctionhouse"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedAuctionHouseServer
	local_timestamp int64
	auction         Auction
}

type Auction struct {
	highest_Bid int64
	isFinished  bool
	itemName    string
	//timeLimit   int64
}

func GetTimestamp(s *Server, i int64) int64 {
	l := float64(s.local_timestamp)
	i_ := float64(i)
	f := math.Max(l, i_) + 1
	return int64(f)
}

func logToFile() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(fmt.Sprint("auctions.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

/*
func setTime (seconds int64) {
	i := int64
	while seconds > i {
		i++
	}
	time.Sleep(1*time.Second)
}
*/

func (s *Server) Bid(ctx context.Context, amount *proto.Amount) (*proto.Ack, error) { //Add timestamp check
	if amount.GetBid() < s.auction.highest_Bid {
		return &proto.Ack{Fail: "You have been outbidded"}, nil //Add highest bid later
	}

	s.auction.highest_Bid = amount.GetBid()
	return &proto.Ack{Success: "Congratulations! You now have the highest Bid!"}, nil //add value later

}

func (s *Server) Result(ctx context.Context, dummyMsg *proto.Void) (*proto.Outcome, error) {
	if s.auction.isFinished {
		fmt.Printf("Auction has Finished. Winning bid was %v", s.auction.highest_Bid)
		return &proto.Outcome{HighestBid: s.auction.highest_Bid}, nil
	} else {
		fmt.Printf("Auction is still ongoing. Highest Bid is %v", s.auction.highest_Bid)
		return &proto.Outcome{HighestBid: s.auction.highest_Bid}, nil
	}
}

func main() {

	logToFile()
	var ThisAuctionHouseServer proto.UnimplementedAuctionHouseServer

	server := &Server{ThisAuctionHouseServer, 0, Auction{}}
	//server.local_timestamp++

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	log.Println("Starting server at port :8080")
	fmt.Println("Starting server at port :8080")

	proto.RegisterAuctionHouseServer(grpcServer, server)
	grpcServer.Serve(listener)

	auc := Auction{
		highest_Bid: 0,
		isFinished:  false,
		itemName:    "Iphone 69 Pro Max",
	}

	server.auction = auc

	i := 0
	for i < 90 {
		i++
		time.Sleep(1 * time.Second)
	}
	server.auction.isFinished = true

}
