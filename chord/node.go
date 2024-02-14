package chord

import (
	"dhtchord/utils"
	"fmt"
)

const (
	KS int = 9
	HS int = 1 << KS
)

var Messages []int // used for metrics

type Finger struct {
	Start     int
	Successor *Node
}

type Node struct {
	Ip          string
	Id          int
	Data        map[int][][2]string
	FingerTable []Finger
	Predecessor *Node
}

func NewNode(ip string) Node {
	return Node{ip, utils.Hash(ip), make(map[int][][2]string), make([]Finger, KS), nil}
}

// Finds the Successor of id
func (n *Node) FindSuccessor(id int) *Node {
	predecessor := n.findPredecessor(id)
	return predecessor.FingerTable[0].Successor
}

// Finds the Predecessor of id
func (n *Node) findPredecessor(id int) *Node {
	current := n
	hops := 1 // used for metrics
	for utils.NotInRangeExclusive(id, current.Id, current.FingerTable[0].Successor.Id) {
		current = current.closestPrecedingFinger(id)
		hops += 1 // used for metrics
	}
	Messages = append(Messages, hops) // used for metrics
	return current
}

// Find the closest finger that precedes id
func (n *Node) closestPrecedingFinger(id int) *Node {
	for i := KS - 1; i >= 0; i-- {
		if utils.InRangeExclusive(n.FingerTable[i].Successor.Id, n.Id, id) {
			return n.FingerTable[i].Successor
		}
	}
	return n
}

// Initializes the finger table of node n
func (n *Node) InitFingerTable(bootstrap *Node) {
	for i := 0; i < KS; i++ {
		n.FingerTable[i].Start = (n.Id + utils.Pow(2, i)) % HS
		n.FingerTable[i].Successor = n
	}

	n.FingerTable[0].Successor = bootstrap.FindSuccessor(n.FingerTable[0].Start)
	successor := n.FingerTable[0].Successor
	predecessor := successor.Predecessor
	n.Predecessor = predecessor
	successor.Predecessor = n
	predecessor.FingerTable[0].Successor = n

	for i := 0; i < KS-1; i++ {
		if utils.InRangeExclusive(n.FingerTable[i+1].Start, n.Id, n.FingerTable[i].Successor.Id) {
			n.FingerTable[i+1].Successor = n.FingerTable[i].Successor
		} else {
			n.FingerTable[i+1].Successor = bootstrap.FindSuccessor(n.FingerTable[i+1].Start)
		}
	}
}

// Updates the finger table of previous nodes
func (n *Node) UpdateOthers() {
	for i := 0; i < KS; i++ {
		pred := n.findPredecessor(utils.DistExclusive(n.Id, utils.Pow(2, i)))
		pred.updateFingerTable(n, i)
	}
}

func (n *Node) updateFingerTable(n0 *Node, i int) {
	if utils.InRangeExclusive(n0.Id, n.Id, n.FingerTable[i].Successor.Id) {
		n.FingerTable[i].Successor = n0
		pred := n.Predecessor
		pred.updateFingerTable(n0, i)
	}
}

func (n *Node) FixFingers() {
	for i := 1; i < KS; i++ {
		n.FingerTable[i].Successor = n.FindSuccessor(n.FingerTable[i].Start)
	}
}

// When joining, import from the successor the keys that node n is responsible for
func (n *Node) FetchData(n0 *Node) {
	if len(n0.Data) > 0 {
		for id := range n0.Data {
			if utils.InRange(id, n.Predecessor.Id, n.Id){
				n.Data[id] = n0.Data[id]
				delete(n0.Data, id)
			}
		}
	}
}

// When leaving, move the keys (data) to successor
func (n *Node) MoveData(n0 *Node) {
	for i := range n.Data{
		n0.Data[i] = n.Data[i]
	}
}

func (n *Node) Leave() {
	successor := n.FingerTable[0].Successor
	predecessor := n.Predecessor
	successor.Predecessor = predecessor
	predecessor.FingerTable[0].Successor = successor
	
	predecessor.FixFingers()
	successor.UpdateOthers()

	n.MoveData(successor)
}

func (n *Node) String() {
	fmt.Printf("Node id: %v | Predecessor id: %v | Successor id: %v\n", n.Id, n.Predecessor.Id, n.FingerTable[0].Successor.Id)
	fmt.Print("Data ids: ")
	for key := range n.Data {
        fmt.Print(key, " ")
    }
	fmt.Println("")
	fmt.Println("")
}