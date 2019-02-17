package blockchain

import (
	"context"

	pbBlock "github.com/prakashpandey/tinycoin-go/proto/block"
)

// LatestBlock returns latest block of blockchain
func (b *Blockchain) LatestBlock(c context.Context, null *pbBlock.Null) (*pbBlock.Block, error) {
	latestBlock := b.GetLastBlock()
	return &pbBlock.Block{
		Index:     int64(latestBlock.Index),
		Timestamp: latestBlock.Timestamp,
		Data:      latestBlock.Data,
		PrevHash:  latestBlock.PrevHash,
		Hash:      latestBlock.Hash,
	}, nil
}
