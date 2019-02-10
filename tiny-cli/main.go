package main

import (
	"fmt"
	"github.com/prakashpandey/gochain/tiny-cli/conf"
)

var flags *conf.Flags

func main() {
	flags = conf.NewFlags()
	flags.Define()
	flags.Parse()
	fmt.Println(flags.String())
}
