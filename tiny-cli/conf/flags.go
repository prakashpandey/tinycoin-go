package conf

import (
	"flag"
	"fmt"
)

type Flags struct {
	IP   string
	Port int
}

func NewFlags() *Flags {
	return &Flags{}
}

func (f *Flags) Define() {
	flag.StringVar(&f.IP, "ip", "localhost", "enter server ip")
	flag.IntVar(&f.Port, "port", 8181, "enter port")
}

func (f *Flags) Parse() {
	flag.Parse()
}

func (f *Flags) String() string {
	return fmt.Sprintf("IP[%s], Port[%d]", f.IP, f.Port)
}
