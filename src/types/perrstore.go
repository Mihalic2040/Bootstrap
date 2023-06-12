package types

import "github.com/multiformats/go-multiaddr"

type PeerInfo struct {
	Peer string                `json:"peer"`
	Addr []multiaddr.Multiaddr `json:"addr"`
}
