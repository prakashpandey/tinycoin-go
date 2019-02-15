package main

import (
	"github.com/prakashpandey/tinycoin-go/tiny-cli/conf"
	"github.com/prakashpandey/tinycoin-go/tiny-cli/server"
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
