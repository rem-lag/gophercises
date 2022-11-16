package hparse

import "io"

// Link = a href element
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
