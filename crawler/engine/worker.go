package engine

import (
	"log"
	"richie.com/richie/learn_golang/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parser(body,r.Url), nil
}
