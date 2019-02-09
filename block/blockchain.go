package block

import (
	"encoding/json"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{GenesisBlock()},
	}
}

func (bc *Blockchain) Add(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) String() string {
	if blockchain, err := json.MarshalIndent(bc.Blocks, "", "\t"); err != nil {
		return ""
	} else {
		return string(blockchain)
	}
}
