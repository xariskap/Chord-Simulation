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
var ring = chord.NewChord()

func BenchmarkJoin(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			build(ipArray[:v.input], &ring)
		})
	}
}
