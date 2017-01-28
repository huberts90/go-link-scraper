# go-link-scraper

Simple link scaper tested with various websites.

## Installation

go get github.com/huberts90/go-link-scraper

## Usage

package main

import (
  "github.com/huberts90/scraper"
  "os"
  "fmt"
)

func main() {
	maxWorkers := 3
  chUrls := make(chan string, 10)
	chResult := make(chan []byte, 10)
	chFinished := make(chan bool)

  seedUrls := os.Args[1:]
	linksJSON := make([]byte, 0)

	countUrls := len(seedUrls)
	for i := 1; i <= scraper.Min(countUrls, maxWorkers); i++ {
			go scraper.Crawl(i, chUrls, chResult, chFinished)
	}

	for _, url := range seedUrls {
		chUrls <- url
	}
	close(chUrls)

	for c := 0; c < countUrls; {
		select {
			case linksJSON = <- chResult:
				fmt.Printf("JSON %s", linksJSON);
			case <-chFinished:
				c++
		}
	}
}
