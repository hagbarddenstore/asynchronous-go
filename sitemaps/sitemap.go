package sitemaps

import (
	"bytes"
	"sort"
	"strings"
)

// Sitemap holds information about an entire site.
type Sitemap struct {
	RootURL string
	Pages   []*Page
}

// Page holds information about a single page and it's child pages.
type Page struct {
	URL   string
	Pages []*Page
}

// NewSitemap initializes a new Sitemap.
func NewSitemap(rootURL string, links []string) *Sitemap {
	sort.Strings(links)

	var pages []*Page

	for _, link := range links {
		parts := getParts(link)

		insertItem(&pages, parts)
	}

	return &Sitemap{RootURL: rootURL, Pages: pages}
}

func getParts(link string) []string {
	parts := strings.Split(link, "/")

	var items []string

	for _, part := range parts {
		if part != "" {
			items = append(items, part)
		}
	}

	return items
}

func insertItem(pages *[]*Page, item []string) {
	if len(item) == 0 {
		return
	}

	var foundPage *Page

	for _, page := range *pages {
		if page.URL == item[0] {
			foundPage = page

			break
		}
	}

	if foundPage == nil {
		foundPage = &Page{
			URL:   item[0],
			Pages: make([]*Page, 0),
		}

		*pages = append(*pages, foundPage)
	}

	if len(item) > 1 {
		insertItem(&foundPage.Pages, item[1:])
	}
}

func (s *Sitemap) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(s.RootURL + "\n")

	buffer.WriteString(s.printPages(s.Pages, ""))

	return buffer.String()
}

func (s *Sitemap) printPages(pages []*Page, indent string) string {
	var buffer bytes.Buffer

	for index, page := range pages {
		if index+1 == len(pages) {
			buffer.WriteString(indent + "└── " + page.URL + "\n")

			if len(page.Pages) > 0 {
				buffer.WriteString(s.printPages(page.Pages, indent+"    "))
			}
		} else {
			buffer.WriteString(indent + "├── " + page.URL + "\n")

			if len(page.Pages) > 0 {
				buffer.WriteString(s.printPages(page.Pages, indent+"│   "))
			}
		}
	}

	return buffer.String()
}
