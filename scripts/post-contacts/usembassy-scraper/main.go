// This script reads the https://www.usembassy.gov/ website, which lists
// contact information for all posts, and pulls all the contacts into a
// single CSV file.
package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Country struct {
	name string
	url  string
}

type Post struct {
	country string
	contact string
	name    string
	website string
}

// getAllCountries reads the main embassy page to get a list
// of contact pages listing the missions in each country.
func getAllCountries() []Country {
	c := initScraper(false)

	var countries []Country

	c.OnHTML("li.listed-post", func(e *colly.HTMLElement) {
		country := Country{}

		country.url = e.ChildAttr("a", "href")
		country.name = e.ChildAttr("a", "title")

		countries = append(countries, country)
	})

	c.Visit("https://www.usembassy.gov")

	fmt.Printf("Found %d countries", len(countries))

	return countries
}

// getContactInfo iterates through each country, accessing it's contact
// page and scraping the data for each mission therein.
func getContactInfo(countries []Country) []Post {
	var posts []Post
	var countryName string

	c := initScraper(true)

	c.Limit(&colly.LimitRule{
		Parallelism: 5,
	})

	c.OnHTML("header h1", func(e *colly.HTMLElement) {
		countryName = strings.TrimSpace(e.Text)
	})

	// The contact pages have two different formats - some have each post
	// in it's owen div with the class `cityname1`. Others have a single
	// `p` tag containing all the information for a single mission. For
	// this reason we require two cases to search on.

	// Case one - single post in `p` tag.
	c.OnHTML("div.entry-content", func(e *colly.HTMLElement) {

		// If the cityname div exists skip to the next case.
		if e.DOM.Has("div.cityname1").Length() > 0 {
			return
		} else {
			post := Post{}

			post.country = countryName
			post.name = e.ChildText("strong")
			post.website = e.ChildAttr("a", "href")
			post.contact = e.DOM.Find("p").First().Text()

			posts = append(posts, post)
		}
	})

	// Case two - each post in `cityname1` div.
	c.OnHTML("div.cityname1", func(e *colly.HTMLElement) {
		post := Post{}

		post.country = countryName
		post.name = e.ChildText("strong")
		post.website = e.ChildAttr("a", "href")
		post.contact = e.Text

		posts = append(posts, post)
	})

	for _, country := range countries {
		c.Visit(country.url)
	}

	c.Wait()

	return posts
}

func main() {
	countries := getAllCountries()
	contacts := getContactInfo(countries)

	writeToCSV(contacts)
}
