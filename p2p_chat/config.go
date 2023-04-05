package main

import "github.com/multiformats/go-multiaddr"

type Config struct {
	NodeName       string
	NetworkId      string
	BootstrapPeers []multiaddr.Multiaddr
	ProtocolID     string
	MaxPeers       int
}

var DefaultConfig = &Config{
	NodeName:  "Netko",
	NetworkId: "bc team",
	BootstrapPeers: []multiaddr.Multiaddr{
		StringsToAddr("/ip4/10.100.2.134/tcp/51425/ipfs/QmYpySQsLoN51x35ndQv5BbHwUzUcd6ZHxiZ6ZauHDCEGc"),
	},
	ProtocolID: "/bc_p2p/1.1.0",
	MaxPeers:   25,
}

func StringsToAddr(addrString string) multiaddr.Multiaddr {
	addr, _ := multiaddr.NewMultiaddr(addrString)
	return addr
}
