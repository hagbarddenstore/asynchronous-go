package scrapers

import (
	"errors"
	"strings"
)

// AsyncSiteScraper scrapes an entire site asynchronously.
type AsyncSiteScraper struct {
	Scraper
	scrapedURLs []string
}

// NewAsyncSiteScraper initializes a new AsyncSiteScraper.
func NewAsyncSiteScraper(siteURL string) (*AsyncSiteScraper, error) {
	siteURL, err := parseSiteURL(siteURL)

	if err != nil {
		return nil, err
	}

	scraper := &AsyncSiteScraper{
		Scraper: Scraper{
			siteURL: siteURL,
			links:   make([]string, 0),
		},
		scrapedURLs: make([]string, 0),
	}

	return scraper, nil
}

// Scrape the site for links.
func (s *AsyncSiteScraper) Scrape() error {
	return errors.New(`Not implemented yet!`)
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
