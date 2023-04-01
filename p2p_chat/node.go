package main

import (
	"context"
	"github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/network"
)

var nLogger = log.Logger("node")


type Node struct {
	*Config
	context.Context

	discovery   *discovery

	inputStream chan string
	outputStream chan *Message

	peers 		*peerSet

}

func NewNode(c *Config, inputStream chan string, outputStream chan *Message) *Node {
	n := &Node {
		Config:  		c,
		Context: 		context.Background(),
		inputStream:	inputStream,
		outputStream:	outputStream,
		peers:			newPeerSet(),
	}

	n.discovery = NewDiscovery(n, n.Config)

	return n
}



func (n *Node) handleNewPeer(stream network.Stream)  {

	if n.peers.Len() < n.MaxPeers {

		peer := NewPeer(stream)
		err := n.peers.Register(peer)
		if err != nil {
			peer.Close()
			return
		}

		nLogger.Info("Connected to new peer", peer.ID())

		go n.readData(peer)
	} else {
		stream.Close()
	}



}

func (n *Node) readData(peer *Peer)  {
	rw := peer.rw
	for {
		m, err := Deserialize(rw)
		if err != nil {
			n.peers.Unregister(peer.ID())
			// break this loop (and goroutine) if there is an error when communicating with this peer
			return
		}

		n.Broadcast(m)
	}
}

func (n *Node) startDiscovery()  {
	n.discovery.Start(n.handleNewPeer)
}


func (n *Node) Start()  {
	go n.startDiscovery()
	go n.writeMessages()
}

func (n *Node) writeMessages()  {

	for {
		data := n.readFromInput()


		m := &Message {
			Sender: n.NodeName,
			Msg:	data,
		}

		go n.Broadcast(m)
	}

}

func (n *Node) Broadcast(m *Message)  {

	data := m.Serialize()
	for _, peer := range n.peers.peers {
		rw := peer.rw
		_, err := rw.Write(data)

		if err == nil {
			err = rw.Flush()
		}

		if err != nil {
			// error communicating with this peer, remove him
			n.peers.Unregister(peer.ID())
			// but still send to other peers
			continue
		}
	}

}

func (n *Node) readFromInput() string {
	return <- n.inputStream
}

func (n *Node) sendToOutput(m *Message)  {
	n.outputStream <- m
}