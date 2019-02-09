package main

import (
	"fmt"
	"github.com/prakashpandey/gochain/block"
)

func main() {
	genesisBlock := block.GenesisBlock()
	b := block.New(genesisBlock)
	b.Data = []byte("I am a transaction")
	fmt.Println(b.String())
}
