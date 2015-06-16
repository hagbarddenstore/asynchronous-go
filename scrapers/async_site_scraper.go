package scrapers

import (
	"strings"
)

// AsyncSiteScraper scrapes an entire site asynchronously.
type AsyncSiteScraper struct {
	BaseScraper
	scrapedURLs []string
}

// NewAsyncSiteScraper initializes a new AsyncSiteScraper.
func NewAsyncSiteScraper(siteURL string) (*AsyncSiteScraper, error) {
	siteURL, err := parseSiteURL(siteURL)

	if err != nil {
		return nil, err
	}

	scraper := &AsyncSiteScraper{
		BaseScraper: BaseScraper{
			siteURL: siteURL,
			links:   make([]string, 0),
		},
		scrapedURLs: make([]string, 0),
	}

	return scraper, nil
}

// Scrape the site for links.
func (s *AsyncSiteScraper) Scrape() error {
	urlsToScrape := 1
	scrapedURLs := 0

	uniqueLinks := make(chan []string)
	errors := make(chan error)

	go s.scrapeURL(s.siteURL, uniqueLinks, errors)

	for {
		select {
		case links := <-uniqueLinks:
			scrapedURLs++

			for _, link := range links {
				s.links = append(s.links, link)

				if s.urlIsScraped(link) {
					continue
				}

				s.scrapedURLs = append(s.scrapedURLs, link)

				urlsToScrape++

				go s.scrapeURL(s.prependSiteURL(link), uniqueLinks, errors)
			}

			if urlsToScrape == scrapedURLs {
				return nil
			}

		case err := <-errors:
			return err
		}
	}
}

func (s *AsyncSiteScraper) scrapeURL(
	URL string,
	uniqueLinks chan []string,
	errors chan error) {
	scraper := NewPageScraper(s.siteURL, URL)

	err := scraper.Scrape()

	if err != nil {
		errors <- err

		return
	}

	foundLinks := scraper.UniqueLinks()

	uniqueLinks <- foundLinks
}

func (s *AsyncSiteScraper) urlIsScraped(URL string) bool {
	for _, url := range s.scrapedURLs {
		if url == URL {
			return true
		}
	}

	return false
}

func (s *AsyncSiteScraper) prependSiteURL(URL string) string {
	if !strings.HasPrefix(URL, s.siteURL) {
		return s.siteURL + URL[1:]
	}

	return URL
}

type scrapeResult struct {
	Links []string
}
