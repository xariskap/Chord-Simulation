package chord

import (
	"dhtchord/utils"
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
		n.FetchData(n.FingerTable[0].Successor)
	}

	c.Nodes = append(c.Nodes, n)
}

// Node n leaves the ring and passes the data to the immediate successor
func (c *Chord) Leave(n *Node) {
	n.Leave()

	for i, node := range c.Nodes{
		if node.Id == n.Id{
			c.Nodes = append(c.Nodes[:i], c.Nodes[i+1:]...)
		}
	}
}

// Returns a random node of the network
func (c Chord) bootstrapNode() *Node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(c.Nodes))
	entry := c.Nodes[randomIndex]
	return entry
}

// When joining, import from the successor the keys that node n is responsible for
func (c Chord) ImportData(data []utils.Scientist) {
	var value [2]string
	for _, s := range data {
		id := utils.Hash(s.Education)
		value[0], value[1] = s.Name, s.NumOfAwards
		bootstrap := c.bootstrapNode()
		idSuccessor := bootstrap.FindSuccessor(id)
		idSuccessor.Data[id] = append(idSuccessor.Data[id], value)
	}
}

func (c Chord) Query(edu string){
	id := utils.Hash(edu)
	bootstrap := c.bootstrapNode()
	idSuccessor := bootstrap.FindSuccessor(id)
	fmt.Println(idSuccessor.Data[id])
}

func (c Chord) Lookup(id int) bool{
	bootstrap := c.bootstrapNode()
	idSuccessor := bootstrap.FindSuccessor(id)
	return idSuccessor.Id == id
}

func (c Chord) String() {
	for _, v := range c.Nodes {
		fmt.Printf("Node: %v, Predecessor %v, Successor %v \n", v.Id, v.Predecessor.Id, v.FingerTable[0].Successor.Id)
	}
}

func (ring Chord) Demo() {
	ring.String()
	fmt.Println("")
	fmt.Println("Printing node ", ring.Nodes[8].Id)
	ring.Nodes[8].String()
	fmt.Println("Printing node ", ring.Nodes[9].Id)
	ring.Nodes[9].String()
	fmt.Println("---------------------------------------------")
	fmt.Println("Node", ring.Nodes[8].Id, "leaves the ring...")
	id := ring.Nodes[8].Id
	ring.Leave(ring.Nodes[8])
	fmt.Println("Lookup for node", id,":", ring.Lookup(id))
	fmt.Println("")
	ring.String()
	fmt.Println("")
	fmt.Println("Printing node ", ring.Nodes[8].Id)
	ring.Nodes[8].String()
	fmt.Println("---------------------------------------------")
	node := NewNode("10.10.20.30:5432")
	ring.Join(&node)
	fmt.Println("Node", ring.Nodes[9].Id, "enters the ring...")
	fmt.Println("Lookup for node", id,":", ring.Lookup(id))
	fmt.Println("Printing node ", ring.Nodes[9].Id)
	ring.Nodes[9].String()
	fmt.Println("Printing node ", ring.Nodes[8].Id)
	ring.Nodes[8].String()
	fmt.Println("")
	ring.String()
}
