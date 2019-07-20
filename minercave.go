package main

import (
	"github.com/spatocode/minercave/app"
	"github.com/spatocode/minercave/net"
)

var cfg net.Config

func main() {
	app.Configure(&cfg)
	app.Exec(&cfg)
}