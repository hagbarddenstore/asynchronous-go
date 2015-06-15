package scrapers

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
)

// Parser parses a HTML page and stores the found links.
type Parser struct {
	reader io.Reader
	links  []string
}

// NewParser initializes a new Parser.
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		reader: reader,
		links:  make([]string, 0),
	}
}

// Parse a HTML page and store the found links.
func (p *Parser) Parse() error {
	tokenizer := html.NewTokenizer(p.reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()

			if err == io.EOF {
				return nil
			}

			return err

		case html.StartTagToken:
			token := tokenizer.Token()

			if token.DataAtom == atom.A {
				for _, attribute := range token.Attr {
					if attribute.Key == "href" {
						link := attribute.Val

						p.links = append(p.links, link)
					}
				}
			}
		}
	}
}

// Links returns the found links.
func (p *Parser) Links() []string {
	return p.links
}
