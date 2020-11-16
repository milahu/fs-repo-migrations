package config

import "github.com/ipfs/fs-repo-migrations/ipfs-10-to-11/_vendor/github.com/libp2p/go-libp2p-core/peer"

// Peering configures the peering service.
type Peering struct {
	// Peers lists the nodes to attempt to stay connected with.
	Peers []peer.AddrInfo
}
