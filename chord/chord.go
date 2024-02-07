package chord

import (
	"fmt"
	"math/rand"
	"time"
)

type Chord struct {
	Nodes []*Node
}

func NewChord() Chord {
	return Chord{make([]*Node, 0)}
}

// Node n joins the ring
func (c *Chord) Join(n *Node) {
	if len(c.Nodes) == 0 {
		n.Predecessor = n
		n.InitFingerTable(n)
	} else {
		bootstrap := c.bootstrapNode()
		n.InitFingerTable(bootstrap)
		n.UpdateOthers()
	}

	c.Nodes = append(c.Nodes, n)
}

// returns a random node of the network
func (c Chord) bootstrapNode() *Node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(c.Nodes))
	entry := c.Nodes[randomIndex]
	return entry
}

func (c Chord) String() {
	for _, v := range c.Nodes {
		fmt.Printf("Node: %v, Predecessor %v, Successor %v \n", v.Id, v.Predecessor.Id, v.FingerTable[0].Successor.Id)
	}
}
