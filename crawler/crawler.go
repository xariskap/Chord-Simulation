package crawler

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"

	"github.com/gocolly/colly"
)

type Scientist struct {
	Education   string `json:"education"`
	Name        string `json:"name"`
	NumOfAwards string `json:"awards"`
}

var scientists []Scientist

func Crawl() {

	numOfAwards := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	education := []string{
		"MIT",
		"CEID",
		"Harvard",
		"Stanford",
		"Caltech",
		"ETH Zurich",
		"UCL",
		"Imperial",
		"LSE",
		"NUS",
		"Princeton",
		"Yale",
		"Johns Hopkins",
		"UC Berkeley",
		"UCLA",
		"University of Hong Kong",
		"UPenn",
		"University of Tokyo",
		"UIUC",
		"University of Washington",
		"UT Austin",
		"UBC",
		"UCSD",
		"NYU",
		"University of Michigan",
		"University of Cambridge",
		"University of Oxford",
		"Columbia",
		"University of Toronto",
		"University of Edinburgh",
		"London Business School",
		"INSEAD",
		"HEC Paris",
		"University of Chicago",
		"UC San Francisco",
		"University of Sydney",
		"Carnegie Mellon",
		"EPFL",
		"ETH Zurich",
		"TU Munich",
		"KAIST",
		"University of Melbourne",
		"Peking University",
		"Tsinghua University",
		"Australian National University",
		"University of Copenhagen",
		"University of Amsterdam",
		"University of Stockholm",
		"Seoul National University",
		"University of Vienna",
		"University of Brussels",
		"Moscow State University",
		"Bosphorus University",
		"American University in Dubai",
		"IESE Business School",
		"University of Cape Town",
		"University of Nairobi",
		"Tel Aviv University",
		"University of Jordan",
		"University of Marrakech",
		"University of Casablanca",
		"University of Tunis",
		"Kuwait University",
		"Qatar University",
		"University of Bahrain",
		"Sultan Qaboos University",
		"University of Tehran",
		"University of Baghdad",
		"King Saud University",
		"United Arab Emirates University",
		"University of Havana",
		"University of Puerto Rico",
		"University of Panama",
		"University of the Andes",
		"University of Lima",
		"University of Quito",
		"Central University of Venezuela",
		"National University of Asuncion",
		"University of Montevideo",
		"University of Guyana",
	  }
	scrapeURL := "https://en.wikipedia.org/wiki/List_of_computer_scientists"
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML(".mw-parser-output li", func(e *colly.HTMLElement) {
		edu := education[rand.Intn(len(education))]
		award := numOfAwards[rand.Intn(len(numOfAwards))]
		name := e.DOM.Find("a").First().Text()

		scientists = append(scientists, Scientist{edu, name, award})

	})

	c.Visit(scrapeURL)
	scientists = scientists[:len(scientists)-13]

	file, err := json.MarshalIndent(scientists, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile("data/dataset.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
