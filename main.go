package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
)

func main() {

	ipArray := utils.Parse()
	ring := chord.NewChord()

	// Building the network
	for _, ip := range ipArray {
		node := chord.NewNode(ip)
		ring.Join(&node)
	}

	ring.String()

}
