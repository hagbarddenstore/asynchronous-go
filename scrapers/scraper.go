package scrapers

// Scraper defines a site/page scraper.
type Scraper interface {
	Scrape() error
	UniqueLinks() []string
}
