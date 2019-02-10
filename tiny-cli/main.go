package main

import (
	"github.com/prakashpandey/gochain/tiny-cli/conf"
	"github.com/prakashpandey/gochain/tiny-cli/server"
)

var flags *conf.Flags

func main() {
	flags = conf.NewFlags()
	flags.Define()
	flags.Parse()
	if flags.Connect {
		server.Connect(flags.IP, flags.Port)
	}
}
