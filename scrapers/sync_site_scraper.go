package scrapers

import (
	"strings"
)

// SyncSiteScraper scrapes an entire site synchronously.
type SyncSiteScraper struct {
	Scraper
	scrapedURLs []string
}

// NewSyncSiteScraper initializes a new SyncSiteScraper.
func NewSyncSiteScraper(siteURL string) (*SyncSiteScraper, error) {
	siteURL, err := parseSiteURL(siteURL)

	if err != nil {
		return nil, err
	}

	scraper := &SyncSiteScraper{
		Scraper: Scraper{
			siteURL: siteURL,
			links:   make([]string, 0),
		},
		scrapedURLs: make([]string, 0),
	}

	return scraper, nil
}

// Scrape the site for links.
func (s *SyncSiteScraper) Scrape() error {
	links, err := s.scrapeURL(s.siteURL)

	if err != nil {
		return err
	}

	s.links = links

	return nil
}

func (s *SyncSiteScraper) scrapeURL(URL string) ([]string, error) {
	if s.urlIsScraped(URL) {
		return make([]string, 0), nil
	}

	s.scrapedURLs = append(s.scrapedURLs, URL)

	var links []string

	scraper := NewPageScraper(s.siteURL, URL)

	err := scraper.Scrape()

	if err != nil {
		return nil, err
	}

	foundLinks := scraper.UniqueLinks()

	for _, link := range foundLinks {
		links = append(links, link)

		childLinks, err := s.scrapeURL(s.prependSiteURL(link))

		if err != nil {
			return nil, err
		}

		for _, childLink := range childLinks {
			links = append(links, childLink)
		}
	}

	return links, nil
}

func (s *SyncSiteScraper) urlIsScraped(URL string) bool {
	for _, url := range s.scrapedURLs {
		if url == URL {
			return true
		}
	}

	return false
}

func (s *SyncSiteScraper) prependSiteURL(URL string) string {
	if !strings.HasPrefix(URL, s.siteURL) {
		return s.siteURL + URL[1:]
	}

	return URL
}
