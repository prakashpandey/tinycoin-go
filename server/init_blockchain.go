package main

import (
	"fmt"
	"time"

	"github.com/prakashpandey/tinycoin-go/server/block"
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

func startBlockChain() {
	for i := 0; i < 100000; i++ {
		generateNewBlock()
		fmt.Printf("%s\n", blockchain.String())
		time.Sleep(1 * time.Second)
	}
}
