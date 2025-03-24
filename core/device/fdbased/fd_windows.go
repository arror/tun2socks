package fdbased

import (
	"errors"

	"vpn/core/device"
)

func Open(name string, mtu uint32, offset int) (device.Device, error) {
	return nil, errors.ErrUnsupported
}
