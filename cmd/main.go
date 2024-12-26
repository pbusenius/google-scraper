package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type content struct {
	Results []result `json:"results"`
}

type result struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	Query       string `json:"query"`
}

func export(content content, file string) error {
	rankingsJson, err := json.Marshal(content)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, rankingsJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var results []result
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.google.de/search?q=k25+augsburg")

	results = append(results, result{Url: "test", Description: "test", Query: "test"})
	var content = content{Results: results}

	fmt.Println(content)

	err := export(content, "test.json")
	if err != nil {
		fmt.Println(err)
	}
}
