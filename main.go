package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
)

func main() {

	ipArray := utils.Parse()
	var nodes []*chord.Node

	for _, ip := range ipArray {
		n := chord.NewNode(ip)
		nodes = append(nodes, &n)
	}

	ring := chord.NewChord()

	for _, node := range nodes {
		ring.Join(node)
	}

	for _, v := range ring.Nodes {
		v.String()
	}

}
