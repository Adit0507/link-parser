package link

import "io"

// represents a link in html doc
type Link struct {
	Href string
	Text string
}

// Parse will take in a html doc and return a slice of links
func Parse(r io.Reader) ([]Link, error) {
	
	return nil, nil
}