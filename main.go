package main

import (
	"github.com/beom-yun/sbcoin/blockchain"
	"github.com/beom-yun/sbcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
