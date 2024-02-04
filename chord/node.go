package chord

import (
	"dhtchord/utils"
	"fmt"
	"math"
)

const (
	KS int = 9
	HS int = 1 << KS
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func dist(start, end int) int {
	if start <= end {
		return end - start
	} else {
		return end + HS - start
	}
}

// returns True if key is in (start, end]
func InRange(key, start, end int) bool {
	return dist(start, end) > dist(key, end)
}

type Finger struct {
	Start     int
	Successor *Node
}

type Node struct {
	Ip          string
	Id          int
	Data        map[int][]string
	FingerTable []Finger
	Predecessor *Node
}

func NewNode(ip string) Node {
	return Node{ip, utils.Hash(ip), make(map[int][]string), make([]Finger, KS), nil}
}

func (n *Node) InitFingerTable() {
	for i := 0; i < KS; i++ {
		num := (n.Id + pow(2, i)) % HS
		n.FingerTable[i] = Finger{num, n}
	}
}

func (n *Node) FindSuccessor(id int) *Node {
	if n.Id == n.FingerTable[0].Successor.Id {
		return n
	}

	if n.Predecessor != nil && InRange(id, n.Predecessor.Id, n.Id) {
		return n
	}

	if InRange(id, n.Id, n.FingerTable[0].Successor.Id) {
		return n.FingerTable[0].Successor
	} else {
		n0 := n.ClosestPrecedingNode(id)
		return n0.FindSuccessor(id)
	}
}

func (n *Node) ClosestPrecedingNode(id int) *Node {
	for i := KS - 1; i >= 0; i-- {
		if InRange(n.FingerTable[i].Successor.Id, n.Id, id) {
			return n.FingerTable[i].Successor
		}
	}
	return n
}

func (n *Node) Notify() {
	successor := n.FingerTable[0].Successor
	successor.Predecessor = n
}

func (n *Node) Stabilize() {
	successor := n.FingerTable[0].Successor
	n.Predecessor = successor.Predecessor
}

func (n *Node) FixFingers() {
	for i := 1; i < KS; i++ {
		n.FingerTable[i].Successor = n.FindSuccessor(n.FingerTable[i].Start)
	}
}

func (n *Node) String() {
	fmt.Printf("Node: %v, Predecessor %v, Successor %v \n", n.Id, n.Predecessor.Id, n.FingerTable[0].Successor.Id)
}
