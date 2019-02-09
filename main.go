package main

import (
	"fmt"
	"github.com/prakashpandey/gochain/block"
	"time"
)

// Initialize a new blockchain by creating a
// genesis block
var genesisBlock = block.GenesisBlock()
var blockchain = block.NewBlockchain(genesisBlock)
var prevBlock = genesisBlock

func generateNewBlock() {
	// create new block
	newBlock := block.NewBlock(prevBlock)
	newBlock.Hash = newBlock.CalHash()

	// add new block to blockchain
	blockchain.Add(newBlock)

	prevBlock = newBlock
}

func main() {
	for i := 0; i < 100000; i++ {
		generateNewBlock()
		fmt.Printf("%s\n", blockchain.String())
		time.Sleep(1 * time.Second)
	}
}
