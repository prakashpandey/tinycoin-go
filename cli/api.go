package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pbBlock "github.com/prakashpandey/tinycoin-go/proto/block"
	"google.golang.org/grpc"
)

func getLatestBlock(ip string, port int) {
	address := ip + ":" + strconv.Itoa(port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbBlock.NewBlockServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	block, err := c.LatestBlock(ctx, &pbBlock.Null{})
	if err != nil {
		log.Fatalf("could not get lastest block: %v", err)
	}
	fmt.Printf("%s\n", block.String())
}
