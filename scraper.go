package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type Link struct {
	URL string `json:"url"`
	Text string `json:"text"`
}

type Links struct {
	Links []Link
}

func Crawl(id int, jobs <-chan string, results chan<- []byte, finished chan<- bool) {
	for url := range jobs {
		fmt.Println("Worker ", id, " get to search ", url)
		resp, err := http.Get(url)

	  if err != nil {
	    fmt.Printf("error: failed to crawl %s \n", url)
	    results <- nil
	  }

	  defer resp.Body.Close() // Close body when function ends
	  b, err := ioutil.ReadAll(resp.Body);
		if err != nil {
			fmt.Printf("error: during reading content from url %s \n", url)
	    results <- nil
		}

		results <- findLinks(b)
	}
	// When function finished
	defer func() {
		finished <- true
	}()
}

func findLinks(body []byte) ([]byte) {
  HTMLLinks := Links{}
	link := Link{}
	var buffer bytes.Buffer
  z := html.NewTokenizer(bytes.NewReader(body))
  depth := 0

  for {
      tt := z.Next()
      switch tt {
      case html.ErrorToken:
        	// End of the document, we're done
					// JSON encode
					links, err := json.Marshal(HTMLLinks.Links);
					if err != nil {
						return []byte{}
					}
					return links
			case html.TextToken:
				if depth > 0 {
					buffer.Write(z.Text())
				}
      case html.StartTagToken, html.EndTagToken:
				t := z.Token()
				if t.Data == "a" {
					if tt == html.StartTagToken {
						depth++
						// Find href
						for _, a := range t.Attr {
							if a.Key == "href" {
									link.URL = a.Val
									break
							}
						}
					} else {
						if tt == html.EndTagToken {
							link.Text = buffer.String()
							buffer.Reset()
							HTMLLinks.Links = append(HTMLLinks.Links, link)
						}
						depth--
					}
				}
      }
  }
}
