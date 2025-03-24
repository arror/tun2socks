package main

import (
	_ "vpn/dns"
	"vpn/engine"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
}

//export Startup
func Startup(fd, mtu int) (err error) {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	return engine.Start(fd, mtu)
}

//export Shutdown
func Shutdown() {
	engine.Stop()
}
