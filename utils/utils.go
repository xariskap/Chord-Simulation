package utils

import (
	"bufio"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
)

const HS = 1 << 9

type Scientist struct {
	Education   string
	Name        string
	NumOfAwards string
}

func Hash(data string) int {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	hashInt := new(big.Int).SetBytes(hasher.Sum(nil))
	return int(hashInt.Mod(hashInt, big.NewInt(int64(HS))).Int64())
}

func Parse() []string {
	var ipArray []string
	file, _ := os.Open("data/ip.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := scanner.Text()
		ipArray = append(ipArray, ip)
	}

	return ipArray
}

func JsonToStuct(filePath string) []Scientist {

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(0)
	}

	var scientists []Scientist

	err = json.Unmarshal(jsonData, &scientists)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	return scientists
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Dist(start, end int) int {
	if start <= end {
		return end - start
	} else {
		return end + HS - start
	}
}

// returns True if key is in (start, end]
func InRange(key, start, end int) bool {
	return Dist(start, end) > Dist(key, end)
}

func DistExclusive(start, end int) int {
	if start < end {
		return end - start
	} else {
		return end + HS - start
	}
}

// returns True if key is in (start, end)
func InRangeExclusive(key, start, end int) bool {
	return DistExclusive(start, end) > DistExclusive(key, end)
}

func NotInRangeExclusive(key, start, end int) bool {
	return DistExclusive(start, end) < Dist(key, end)
}
