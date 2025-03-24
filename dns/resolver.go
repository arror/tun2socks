package dns

import (
	"net"

	"vpn/dialer"
)

func init() {
	// We must use this DialContext to query DNS
	// when using net default resolver.
	net.DefaultResolver.PreferGo = true
	net.DefaultResolver.Dial = dialer.DialContext
}
