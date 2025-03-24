package proto

import "fmt"

const (
	Direct Proto = iota
	Reject
	Socks5
)

type Proto uint8

func (proto Proto) String() string {
	switch proto {
	case Direct:
		return "direct"
	case Reject:
		return "reject"
	case Socks5:
		return "socks5"
	default:
		return fmt.Sprintf("proto(%d)", proto)
	}
}
