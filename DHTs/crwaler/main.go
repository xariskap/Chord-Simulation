package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	numOfAwards := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	education := []string{"Massachusetts", "Patras", "Oxford", "Stanford", "Berkeley", "Harvard", "Michigan", "Xidian", "Edinburgh", "Tsinghua", "Carnegie Mellon"}

	// var row [3]string
	data := make([][]string, 0)

	scrapeURL := "https://en.wikipedia.org/wiki/List_of_computer_scientists"
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML(".mw-parser-output li", func(e *colly.HTMLElement) {
		eduRand := education[rand.Intn(len(education))]
		awardRand := numOfAwards[rand.Intn(len(numOfAwards))]
		name := e.DOM.Find("a").First().Text()

		row := []string{name, eduRand, awardRand}
		data = append(data, row)

	})

	c.Visit(scrapeURL)
	data = data[:len(data)-13]

	csvFile, err := os.Create("data.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
