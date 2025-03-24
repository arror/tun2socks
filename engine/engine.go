package engine

import (
	"net/netip"
	"strconv"
	"sync"

	"gvisor.dev/gvisor/pkg/tcpip/stack"

	"vpn/core"
	"vpn/core/device"
	"vpn/core/device/fdbased"
	"vpn/core/option"
	"vpn/proxy"
	"vpn/tunnel"
)

var (
	_mutex  sync.Mutex
	_device device.Device
	_stack  *stack.Stack
)

func Start(fd, mtu int) (err error) {
	_mutex.Lock()
	defer _mutex.Unlock()
	tunnel.T().SetDialer(proxy.NewDirect())
	if _device, err = fdbased.Open(strconv.FormatInt(int64(fd), 10), 1500, 0); err != nil {
		return err
	}
	cfg := &core.Config{
		LinkEndpoint:     _device,
		TransportHandler: tunnel.T(),
		MulticastGroups:  []netip.Addr{},
		Options:          []option.Option{},
	}
	if _stack, err = core.CreateStack(cfg); err != nil {
		return err
	}
	return nil
}

func Stop() {
	_mutex.Lock()
	if _device != nil {
		_device.Close()
	}
	if _stack != nil {
		_stack.Close()
		_stack.Wait()
	}
	_mutex.Unlock()
}
