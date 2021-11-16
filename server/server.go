package main

import (
	"context"
	"math"

	proto "github.com/JoakimGD/MiniProject3/auctionhouse"
)

type Server struct {
	proto.UnimplementedAuctionHouseServer
	local_timestamp int64
	auction         Auction
}

type Auction struct {
	highest_Bid int64
}

func GetTimestamp(s *Server, i int64) int64 {
	l := float64(s.local_timestamp)
	i_ := float64(i)
	f := math.Max(l, i_) + 1
	return int64(f)
}

func (s *Server) Bid(ctx context.Context, amount *proto.Amount) (*proto.Ack, error) {
	if amount.getBid() < s.auction.highest_Bid {
		return
	}
}

func (s *Server) Result(ctx context.Context, amount *proto.Amount) (*proto.Outcome, error) {

}
