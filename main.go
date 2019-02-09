package main

import (
	"fmt"
	"github.com/prakashpandey/gochain/block"
)

var genesisBlock = block.GenesisBlock()
var blockchain = block.NewBlockchain(genesisBlock)

func main() {
	newBlock := block.NewBlock(genesisBlock)
	blockchain.Add(newBlock)
	fmt.Printf("%s\n", blockchain.String())
}
