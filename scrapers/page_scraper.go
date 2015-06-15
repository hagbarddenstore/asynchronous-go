package scrapers

import (
	"net/http"
)

// PageScraper scrapes a website and finds all pages.
type PageScraper struct {
	Scraper
	pageURL string
}

// NewPageScraper initializes a PageScraper.
func NewPageScraper(siteURL string, pageURL string) *PageScraper {
	return &PageScraper{
		Scraper: Scraper{
			siteURL: siteURL,
			links:   make([]string, 0),
		},
		pageURL: pageURL,
	}
}

// Scrape the page for links.
func (s *PageScraper) Scrape() error {
	response, err := http.Get(s.pageURL)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	parser := NewParser(response.Body)

	err = parser.Parse()

	if err != nil {
		return err
	}

	s.links = parser.Links()

	return nil
}
