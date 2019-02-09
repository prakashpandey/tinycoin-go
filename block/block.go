package block

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      []byte `json:"data"`
	PrevHash  []byte `json:"prevHash"`
	// #WARNING #REVISIT 'omitempty' is used so that
	// while hasing it will be ignored
	// assuming that its field will be empty.
	Hash []byte `json:"hash,omitempty"`
}

func GenesisBlock() *Block {
	b := &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Data:      []byte(""),
		PrevHash:  []byte(""),
		Hash:      []byte(""),
	}
	b.Hash = b.CalHash()
	return b
}

func New(b *Block) *Block {
	return &Block{
		Index:     b.Index + 1,
		Timestamp: time.Now().Unix(),
		PrevHash:  b.CalHash(),
	}
}

func (b *Block) CalHash() []byte {
	block, err := json.Marshal(b)
	if err != nil {
		block = []byte{}
	}
	h := sha256.New()
	h.Write(block)
	return h.Sum(nil)
}

func (b *Block) String() string {
	block, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(block)
}
