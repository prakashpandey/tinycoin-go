package main

import (
	"log"
	"net"

	pbBlock "github.com/prakashpandey/tinycoin-go/proto/block"
	"github.com/prakashpandey/tinycoin-go/server/blockchain"
	"google.golang.org/grpc"
)

const (
	port = ":8182"
)

func main() {
	go blockchain.StartBlockChain()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbBlock.RegisterBlockServiceServer(s, blockchain.MainBlockchain)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
