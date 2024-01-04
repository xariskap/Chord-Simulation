package utils

import (
	"bufio"
	"crypto/sha1"
	"math/big"
	"os"
)

const HS = 512

func hashFunc(data string) int {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	hashInt := new(big.Int).SetBytes(hasher.Sum(nil))
	return int(hashInt.Mod(hashInt, big.NewInt(int64(HS))).Int64())
}

func Parse() []int {
	var nodes []int
	file, _ := os.Open("data/ip.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		hashed := hashFunc(ip)
		nodes = append(nodes, hashed)
	}

	return nodes
}
