package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
	"fmt"
)

func main() {

	ipArray := utils.Parse()
	ring := chord.NewChord()

	// Building the network
	for _, ip := range ipArray {
		node := chord.NewNode(ip)
		ring.Join(&node)
	}

	for _, node := range(ring.Nodes){
		node.FixFingers()
	}

	ring.String()

	for _,v := range(ring.Nodes[6].FingerTable){
		fmt.Println(v.Start, v.Successor.Id)
	}
}
