package blockchain

import (
	"fmt"
	"sync"

	"github.com/beom-yun/sbcoin/db"
	"github.com/beom-yun/sbcoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockChain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
