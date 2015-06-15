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

	if async {
		scraper, err := scrapers.NewAsyncSiteScraper(siteURL)

		if err != nil {
			panic(err)
		}

		err = scraper.Scrape()

		if err != nil {
			panic(err)
		}

		links = scraper.UniqueLinks()
	} else {
		scraper, err := scrapers.NewSyncSiteScraper(siteURL)

		if err != nil {
			panic(err)
		}

		err = scraper.Scrape()

		if err != nil {
			panic(err)
		}

		links = scraper.UniqueLinks()
	}

	sitemap := sitemaps.NewSitemap(siteURL, links)

	fmt.Println(sitemap.String())
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
