package main

import "github.com/multiformats/go-multiaddr"


type Config struct {
	NodeName		 string
	NetworkId		 string
	BootstrapPeers   []multiaddr.Multiaddr
	ProtocolID       string
	MaxPeers 	     int
}



var DefaultConfig = & Config {
	NodeName: 		"Matija",
	NetworkId: 		"bc team",
	BootstrapPeers: []multiaddr.Multiaddr{
		StringsToAddr("/ip4/192.168.43.30/tcp/56586/ipfs/QmYDeDmaZ9UEEQk6ebXB9a89oqFZeyxkQmQvMBFn6PrYJi"),
	},
	ProtocolID: 	"/bc_p2p/1.1.0",
	MaxPeers:  		3,
}

func StringsToAddr(addrString string) multiaddr.Multiaddr {
	addr, _ := multiaddr.NewMultiaddr(addrString)
	return addr
}

