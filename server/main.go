package main

import (
	"context"
	"log"
	"net"

	pbBlock "github.com/prakashpandey/tinycoin-go/proto/block"
	"google.golang.org/grpc"
)

const (
	port = ":8182"
)

type blockHelper struct{}

// LatestBlock returns latest block of blockchain
func (s *blockHelper) LatestBlock(c context.Context, null *pbBlock.Null) (*pbBlock.Block, error) {
	latestBlock := blockchain.GetLastBlock()
	return &pbBlock.Block{
		Index:     int64(latestBlock.Index),
		Timestamp: latestBlock.Timestamp,
		Data:      latestBlock.Data,
		PrevHash:  latestBlock.PrevHash,
		Hash:      latestBlock.Hash,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbBlock.RegisterBlockServiceServer(s, &blockHelper{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
