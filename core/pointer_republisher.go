package core

import (
	"github.com/djali-foundation/djali-go/net/repointer"
)

// StartPointerRepublisher - setup republisher for IPNS
func (n *OpenBazaarNode) StartPointerRepublisher() {
	n.PointerRepublisher = net.NewPointerRepublisher(n.DHT, n.Datastore, n.PushNodes, n.IsModerator)
	go n.PointerRepublisher.Run()
}
