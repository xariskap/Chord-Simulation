package chord

import (
	"math/rand"
	"time"
)

type Chord struct {
	Nodes []*Node
}

func NewChord() Chord {
	return Chord{make([]*Node, 0)}
}

func (c *Chord) Join(n *Node) {
	if len(c.Nodes) == 0 {
		n.Predecessor = n
		for i := 0; i < KS; i++ {
			n.FingerTable[i] = Finger{(n.Id + pow(2, i)) % HS, n}
		}

	} else {
		n.InitFingerTable()
		bootstrap := c.bootstrapNode()
		succ := bootstrap.FindSuccessor(n.Id)
		n.FingerTable[0].Successor = succ
		n.Stabilize()
		n.Notify()
		n.FixFingers()
	}

	c.Nodes = append(c.Nodes, n)
}

func (c Chord) bootstrapNode() *Node {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	entry := c.Nodes[rng.Intn(len(c.Nodes))]
	return entry
}
