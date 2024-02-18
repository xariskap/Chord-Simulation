package main

import (
	"dhtchord/chord"
	"dhtchord/utils"
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 10},
	{input: 40},
	{input: 70},
	{input: 100},
	{input: 150},
}

var ipArray = utils.Parse("data/test_ip.txt")
var c = chord.NewChord()

var ch = chord.NewChord()
var fullRing = GetRing(ipArray, &ch)

func BenchmarkJoin(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%v", v.input), func(b *testing.B) {
			c = chord.NewChord()
			build(ipArray[:v.input], &c)
		})
	}
}

func BenchmarkLookup(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input size %v", v.input), func(b *testing.B) {
			for i := 0; i < v.input; i++ {
				fullRing.Lookup(fullRing.Nodes[i].Id)
			}
		})
	}
}
