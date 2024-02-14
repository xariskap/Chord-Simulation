package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
)

// used for benchmark
func GetRing(ipArray []string, ring *chord.Chord) *chord.Chord {
	build(ipArray, ring)
	return ring
}

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

	// Parse the data
	scientists := utils.JsonToStuct("data/dataset.json")
	// Parse the nodes IPs
	ipArray := utils.Parse("data/ip.txt")
	// Create a new chord
	ring := chord.NewChord()

	// Build the network
	build(ipArray, &ring)
	// Import data to the network
	ring.ImportData(scientists)

	ring.Demo()
	ring.Query("CEID", 4)
}
