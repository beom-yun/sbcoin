package main

import (
	"github.com/beom-yun/sbcoin/explorer"
	"github.com/beom-yun/sbcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
