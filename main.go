package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
)

func build(ipArray []string, ring *chord.Chord) {
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
}

func main() {

	scientists := utils.JsonToStuct("data/dataset.json")
	ipArray := utils.Parse("data/ip.txt")
	ring := chord.NewChord()

	// Building the network
	build(ipArray, &ring)

	ring.ImportData(scientists)

	ring.Demo()
	//ring.Query("CEID", 4)

}
