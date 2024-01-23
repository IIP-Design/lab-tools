package main

import (
	"log"

	"github.com/gocolly/colly"
)

// initScraper creates a new colly collection, which can be used to scrape a given website.
func initScraper(async bool) *colly.Collector {
	c := colly.NewCollector(
		colly.Async(async),
	)

	// Setting a valid User-Agent header help to prevent getting blocked
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	return c
}
