package blockchain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/prakashpandey/tinycoin-go/server/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func NewBlockchain(genesisBlock *block.Block) *Blockchain {
	return &Blockchain{
		Blocks: []*block.Block{genesisBlock},
	}
}

func (bc *Blockchain) Add(b *block.Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) GetLastBlock() *block.Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) String() string {
	if blockchain, err := json.MarshalIndent(bc.Blocks, "", "\t"); err != nil {
		return ""
	} else {
		return string(blockchain)
	}
}

// Initialize a new blockchain by creating a
// genesis block
var genesisBlock = block.GenesisBlock()
var MainBlockchain = NewBlockchain(genesisBlock)
var prevBlock = genesisBlock

func generateNewBlock() {
	// create new block
	newBlock := block.NewBlock(prevBlock)
	newBlock.Hash = newBlock.CalHash()

	// add new block to blockchain
	MainBlockchain.Add(newBlock)

	prevBlock = newBlock
}

func StartBlockChain() {
	for i := 0; i < 100000; i++ {
		generateNewBlock()
		fmt.Printf("%s\n", MainBlockchain.String())
		time.Sleep(1 * time.Second)
	}
}
