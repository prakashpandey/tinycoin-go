package block

import (
	"encoding/json"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain(genesisBlock *Block) *Blockchain {
	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) Add(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) String() string {
	if blockchain, err := json.MarshalIndent(bc.Blocks, "", "\t"); err != nil {
		return ""
	} else {
		return string(blockchain)
	}
}
