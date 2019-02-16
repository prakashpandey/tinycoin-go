package main

import (
	"github.com/prakashpandey/tinycoin-go/cli/conf"
)

var flags *conf.Flags

func main() {
	flags = conf.NewFlags()
	flags.Define()
	flags.Parse()
	if flags.Connect {
		getLatestBlock(flags.IP, flags.Port)
	}
}
