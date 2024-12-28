package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
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
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

	results = append(results, result{Url: "test", Description: "test", Query: "test"})
	var content = content{Results: results}

	fmt.Println(content)

	err = export(content, "test.json")
	if err != nil {
		fmt.Println(err)
	}
}
