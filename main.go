package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
	"fmt"
)

func main() {

	nodeIds := utils.Parse()
	chordnet := chord.NewChord()

	for _, id := range nodeIds {
		chordnet.Join(chord.NewNode(id, &chordnet), 1000)
	}

	// temp := chord.NewNode(2, &chordnet)
	// chordnet.Join(temp)

	fmt.Println(chordnet.Nodes)
	chordnet.String()
}
