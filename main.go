package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
	"fmt"
)

func main() {

	scientists := utils.JsonToStuct("data/dataset.json")
	ipArray := utils.Parse()
	ring := chord.NewChord()

	// Building the network
	for _, ip := range ipArray {
		node := chord.NewNode(ip)
		ring.Join(&node)

		// Periodically fix fingers
		if len(ring.Nodes)%5 == 0 {
			for _, node := range ring.Nodes {
				node.FixFingers()
			}
		}
	}

	ring.ImportData(scientists)

	ring.Nodes[9].String()
	fmt.Println("------------------------------------")
	ring.Nodes[8].String()
	fmt.Println(len(ring.Nodes))
	ring.Leave(ring.Nodes[8])
	fmt.Println("------------------------------------")
	ring.Nodes[8].String()
	fmt.Println(len(ring.Nodes))
}
