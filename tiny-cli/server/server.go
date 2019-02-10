package server

import (
	"fmt"
)

func Connect(ip string, port int) {
	fmt.Printf("connected to %s:%d\n", ip, port)
}
