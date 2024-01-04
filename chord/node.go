package chord

import "math"

const (
	KS int = 9
	HS int = 1 << KS
)

func powInt(x, y int) int {
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

type Node struct {
	Id          int
	Network     *Chord
	Data        map[int][]string
	FingerTable Finger
	Predecessor *Node
}

type Finger struct {
	start      []int
	fingerNode []*Node
}

func NewNode(id int, net *Chord) Node {
	return Node{id, net, make(map[int][]string), NewFinger(), nil}
}

func NewFinger() Finger {
	return Finger{make([]int, 0), make([]*Node, 0)}
}

func (n *Node) ClosestPrecedingNode(id int) *Node {
	current := n
	for i := 0; i < KS; i++ {
		if InRange(current.FingerTable.fingerNode[i].Id, current.Id, id) {
			current = current.FingerTable.fingerNode[i]
		}
	}
	return current
}

func (n *Node) Successor(id int) *Node {
	current := n.ClosestPrecedingNode(id)
	next := current.ClosestPrecedingNode(id)

	for InRange(next.Id, current.Id, id) {
		current = next
		next = current.ClosestPrecedingNode(id)
	}

	if current.Id == id {
		return current
	}
	return current.FingerTable.fingerNode[0]
}

func (n *Node) InsertPredecessor(newNode Node) {
	newNode.FingerTable.start = append(newNode.FingerTable.start, (newNode.Id + 1%HS))
	newNode.FingerTable.fingerNode = append(newNode.FingerTable.fingerNode, n)

	n.Predecessor.FingerTable.fingerNode[0] = &newNode
	newNode.Predecessor = n.Predecessor

	n.Predecessor = &newNode

	newNode.FingerTableInit()

}

func (n *Node) FingerTableInit() {
	i := 1
	for i < KS {
		pos := (n.Id + powInt(2, i)) % HS

		for (i < KS) && InRange(pos, n.Id, n.FingerTable.fingerNode[0].Id) {
			n.FingerTable.start = append(n.FingerTable.start, pos)
			n.FingerTable.fingerNode = append(n.FingerTable.fingerNode, n.FingerTable.fingerNode[0])
			i++
			pos = (n.Id + powInt(2, i)) % HS
		}

		if i == KS {
			break
		}

		n.FingerTable.start = append(n.FingerTable.start, pos)
		n.FingerTable.fingerNode = append(n.FingerTable.fingerNode, n.Successor(pos))
		i++
	}
}

func (n *Node) CalcFurthestPredecessor() int {
	if n.Id >= powInt(2, KS-1) {
		return (n.Id - powInt(2, KS-1))
	}
	return powInt(2, KS) + n.Id - (powInt(2, KS-1))
}

func (n *Node) FixFingers() {
	for i := 0; i < KS-1; i++ {

	}

}

func (n *Node) UpdateFingers(){}
