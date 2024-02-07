package chord

import (
	"dhtchord/utils"
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

func distExclusive(start, end int) int {
	if start < end {
		return end - start
	} else {
		return end + HS - start
	}
}

// returns True if key is in (start, end)
func InRangeExclusive(key, start, end int) bool {
	return distExclusive(start, end) > distExclusive(key, end)
}

func NotInRangeExclusive(key, start, end int) bool {
	return distExclusive(start, end) < dist(key, end)
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


func (n *Node) FindSuccessor(id int) *Node {
	predecessor := n.FindPredecessor(id)
	return predecessor.FingerTable[0].Successor
}

func (n *Node) FindPredecessor(id int) *Node {
	current := n
	for NotInRangeExclusive(id, current.Id, current.FingerTable[0].Successor.Id) {
		current = current.ClosestPrecedingFinger(id)
	}
	return current
}

func (n *Node) ClosestPrecedingFinger(id int) *Node {
	for i:= KS -1; i >= 0; i-- {
		if InRangeExclusive(n.FingerTable[i].Successor.Id, n.Id, id){
			return n.FingerTable[i].Successor
		}  
	}
	return n
}

func (n *Node) InitFingerTable(bootstarp *Node){
	for i:= 0; i < KS; i++ {
		n.FingerTable[i].Start = (n.Id + pow(2, i)) % HS
		n.FingerTable[i].Successor = n
	}

	n.FingerTable[0].Successor = bootstarp.FindSuccessor(n.FingerTable[0].Start)
	successor := n.FingerTable[0].Successor
	predecessor := successor.Predecessor
	n.Predecessor = predecessor
	successor.Predecessor = n
	predecessor.FingerTable[0].Successor = n

	for i:= 0; i < KS - 1; i++ {
		if InRangeExclusive(n.FingerTable[i+1].Start, n.Id, n.FingerTable[i].Successor.Id){
			n.FingerTable[i+1].Successor = n.FingerTable[i].Successor
		} else {
			n.FingerTable[i+1].Successor = bootstarp.FindSuccessor(n.FingerTable[i+1].Start)
		}
	}
}

func (n *Node) UpdateOthers() {
	for i:= 0; i < KS; i++ {
		pred := n.FindPredecessor(distExclusive(n.Id, pow(2, i)))
		pred.UpdateFingerTable(n , i)
	}
}

func (n *Node) UpdateFingerTable(ni *Node, i int) {
	if InRangeExclusive(ni.Id, n.Id, n.FingerTable[i].Successor.Id){
		n.FingerTable[i].Successor = ni
		pred := n.Predecessor
		pred.UpdateFingerTable(ni, i)
	}
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
