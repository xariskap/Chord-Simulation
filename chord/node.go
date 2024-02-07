package chord

import (
	"dhtchord/utils"
	"fmt"
)

const (
	KS int = 9
	HS int = 1 << KS
)

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

func (n *Node) FindSuccessor(id int) *Node {
	predecessor := n.findPredecessor(id)
	return predecessor.FingerTable[0].Successor
}

func (n *Node) findPredecessor(id int) *Node {
	current := n
	for utils.NotInRangeExclusive(id, current.Id, current.FingerTable[0].Successor.Id) {
		current = current.closestPrecedingFinger(id)
	}
	return current
}

func (n *Node) closestPrecedingFinger(id int) *Node {
	for i := KS - 1; i >= 0; i-- {
		if utils.InRangeExclusive(n.FingerTable[i].Successor.Id, n.Id, id) {
			return n.FingerTable[i].Successor
		}
	}
	return n
}

func (n *Node) InitFingerTable(bootstarp *Node) {
	for i := 0; i < KS; i++ {
		n.FingerTable[i].Start = (n.Id + utils.Pow(2, i)) % HS
		n.FingerTable[i].Successor = n
	}

	n.FingerTable[0].Successor = bootstarp.FindSuccessor(n.FingerTable[0].Start)
	successor := n.FingerTable[0].Successor
	predecessor := successor.Predecessor
	n.Predecessor = predecessor
	successor.Predecessor = n
	predecessor.FingerTable[0].Successor = n

	for i := 0; i < KS-1; i++ {
		if utils.InRangeExclusive(n.FingerTable[i+1].Start, n.Id, n.FingerTable[i].Successor.Id) {
			n.FingerTable[i+1].Successor = n.FingerTable[i].Successor
		} else {
			n.FingerTable[i+1].Successor = bootstarp.FindSuccessor(n.FingerTable[i+1].Start)
		}
	}
}

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

func (n *Node) MoveData(n0 *Node) {
	for i := range n.Data{
		n0.Data[i] = n.Data[i]
	}
}

func (n *Node) Leave() {
	sucessor := n.FingerTable[0].Successor
	predecessor := n.Predecessor
	sucessor.Predecessor = predecessor
	predecessor.FingerTable[0].Successor = sucessor
	
	predecessor.FixFingers()
	sucessor.UpdateOthers()

	n.MoveData(sucessor)
}

func (n *Node) String() {
	fmt.Printf("Node id: %v\n", n.Id)
	fmt.Printf("Predecessor: %v\n", n.Predecessor.Id)
	for key, value := range n.Data {
        fmt.Printf("%v: %v\n", key, value)
    }
}