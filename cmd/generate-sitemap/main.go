package main

import (
	"flag"
	"fmt"
	"github.com/hagbarddenstore/asynchronous-go/scrapers"
	"github.com/hagbarddenstore/asynchronous-go/sitemaps"
)

var (
	siteURL     string
	async       bool
	scrapedURLs []string
)

func main() {
	flag.Parse()

	var links []string

	scraper, err := createScraper()

	if err != nil {
		panic(err)
	}

	err = scraper.Scrape()

	if err != nil {
		panic(err)
	}

	links = scraper.UniqueLinks()

	sitemap := sitemaps.NewSitemap(siteURL, links)

	fmt.Println(sitemap.String())
}

func createScraper() (scrapers.Scraper, error) {
	if async {
		return scrapers.NewAsyncSiteScraper(siteURL)
	}

	return scrapers.NewSyncSiteScraper(siteURL)
}

func init() {
	flag.StringVar(
		&siteURL,
		"site-url",
		"http://localhost/",
		"The URL to generate a sitemap for")

	flag.BoolVar(
		&async,
		"async",
		true,
		"Should use the async scraper")
}
