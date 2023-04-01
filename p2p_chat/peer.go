package main

import (
	"bufio"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
)

type Peer struct {
	stream 	network.Stream
	id 		peer.ID

	rw      *bufio.ReadWriter
}

func NewPeer(stream network.Stream) *Peer {

	p := &Peer{
		stream: stream,
		id:		stream.Conn().RemotePeer(),
		rw: 	bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream)),
	}

	return p
}

func (p *Peer) ID() peer.ID {
	return p.id
}

func (p *Peer) Close() error {
	return p.stream.Close()
}

type peerSet struct {
	peers map[peer.ID]*Peer
}

func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[peer.ID]*Peer),
	}
}

func (ps *peerSet) Register(p *Peer) error {

	if _, ok := ps.peers[p.id]; ok {
		return errors.New("already registered")
	}

	ps.peers[p.id] = p

	return nil
}

func (ps *peerSet) Unregister(id peer.ID) error {
	p, ok := ps.peers[id]
	if !ok {
		return errors.New("not registered")
	}

	delete(ps.peers, id)

	return p.Close()
}

func (ps *peerSet) Len() int {
	return len(ps.peers)
}
