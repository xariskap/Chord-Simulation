package chord

import (
	"fmt"
)



type Chord struct {
	Nodes map[int]Node
}

func NewChord() Chord{
	return Chord{make(map[int]Node)}
}

func (c *Chord) Build() {

}


func (chord *Chord) Join(node Node, sn int) {

	// If there are no nodes in the Chord network, set the new node as the first node
	if len(chord.Nodes) == 0 {
		node.Predecessor = &node

		for i:= 0; i < KS; i++{
			node.FingerTable.start = append(node.FingerTable.start, i)
			node.FingerTable.fingerNode = append(node.FingerTable.fingerNode, &node)
		}
		
		chord.Nodes[node.Id] = node
		return
	}

	startingNode := chord.getNode(sn)
	startingNode.Successor(node.Id)

	// Find the predecessor node for the new node

	//predecessor := chord.findPredecessor(node.Id)

	// predecessor := chord.Nodes[1]
	// // // Update successor and predecessor pointers for the new node
	// node.Successor = predecessor.Successor
	// node.Predecessor = &predecessor
	// predecessor.Successor = &node

	// chord.Nodes[node.Id] = node

	// Update finger tables for relevant nodes
	// chord.updateFingerTables(node)

	// chord.nodes[node.id] = node
}

// func (chord *Chord) findPredecessor(id int) *Node {
// 	for _, n := range chord.Nodes {
// 		if compCWDist(n.Id, n.Successor.Id, id) {
// 			return &n
// 		}
// 		pred := n
// 	}
// 	return &c // Should not reach here in a properly initialized Chord network
// }

func (chord *Chord) getNode(id int) Node {
	if _, ok := chord.Nodes[id]; ok{
		return chord.Nodes[id]
	}
	for _, n := range(chord.Nodes){
		return n
	}
	return NewNode(1000, chord)
}

func (chord *Chord) String() {
	for _, v := range(chord.Nodes){
		fmt.Printf("pred: %d -- node: %d -- succ %d\n", v.Predecessor.Id, v.Id, v.FingerTable.fingerNode[0].Id)
	}
}