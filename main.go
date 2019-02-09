package main

import (
	"fmt"
	"github.com/prakashpandey/gochain/block"
)

func main() {
	bc := block.NewBlockchain()
	fmt.Printf("%s\n", bc.String())
}
