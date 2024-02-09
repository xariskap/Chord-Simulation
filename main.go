package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
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

	ring.Demo()
	ring.Query("CEID", 0)
}
