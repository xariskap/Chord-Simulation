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

func (c *Chord) Join(n *Node) {
	
	// c.Initialize(n)
	
	// if len(c.Nodes) >= 2{

	// 	n.InitFingerTable()
	// 	bootstrap := c.bootstrapNode()
	// 	successor := bootstrap.FindSuccessor(n.Id)
	// 	n.FingerTable[0].Successor = successor
	// 	n.Stabilize()
	// 	n.Notify()
	// 	n.FixFingers()
	// }

	// c.Nodes = append(c.Nodes, n)

	if len(c.Nodes) == 0 {
		n.Predecessor = n
		n.Fingers()
		// n.InitFingerTable(n)
	} else {
		bootstrap := c.bootstrapNode()
		n.InitFingerTable(bootstrap)
		//n.UpdateOthers()
	}

	c.Nodes = append(c.Nodes, n)



}

// func (c *Chord) Initialize(n *Node) {
// 	if len(c.Nodes) == 0 {
// 		n.Predecessor = n
// 		n.InitFingerTable()
// 	} else if len(c.Nodes) == 1 {
// 		n.InitFingerTable()
// 		successor := c.Nodes[0]
// 		successor.FingerTable[0].Successor = n
// 		n.FingerTable[0].Successor = successor
// 		n.Stabilize()
// 		n.Notify()
// 		n.FixFingers()
// 		successor.FixFingers()
// 	}
// }

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
