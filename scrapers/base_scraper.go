package scrapers

import (
	"fmt"
	"net/url"
	"strings"
)

// BaseScraper is the base scraper for HTML sites/pages.
type BaseScraper struct {
	siteURL string
	links   []string
}

// UniqueLinks returns all the unique internal links that where found.
func (s *BaseScraper) UniqueLinks() []string {
	var links stringArray

	for _, link := range s.links {
		link = s.removeSiteURL(link)

		if link == "" || link[0] != '/' {
			continue
		}

		if !links.Contains(link) {
			links = append(links, link)
		}
	}

	return links
}

func (s *BaseScraper) removeSiteURL(link string) string {
	if link == "" {
		return link
	}

	if link[0] == '/' {
		return link
	}

	if strings.HasPrefix(link, s.siteURL) {
		link = strings.TrimPrefix(link, s.siteURL)

		if link == "" {
			return "/"
		}

		if link[0] != '/' {
			return "/" + link
		}
	}

	return link
}

func parseSiteURL(URL string) (string, error) {
	u, err := url.Parse(URL)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s/", u.Scheme, u.Host), nil
}
