package scraper

import (
	"testing"
	"bytes"
)

const testHTML = `<html>
									  <body>
									    <div class="bigbird">
									      <div class="container">
									        <div class="bigbird">
									          Hi, I'm Big Bird
														<a href="hubertsiwik.pl">Click on <span>Link 1</span></a>
														<a href="allegro.pl">Link 2</a>
									        </div>
									      </div>
									    </div>
									  </body>
									</html>
									`

// Check whether links are found
func TestFindLinks(t *testing.T) {
		links := findLinks([]byte(testHTML))
		expected := []byte(`[{"url":"hubertsiwik.pl","text":"Click on Link 1"},{"url":"allegro.pl","text":"Link 2"}]`)
		if(!bytes.Equal(links, expected)) {
			t.Error("incorrect json content")
		}
}
