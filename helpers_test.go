package scraper

import (
	"testing"
)

// Check whether links are found
func TestMin(t *testing.T) {
    minValue := Min(2, 4)
		if minValue != 2 {
			t.Error("expected 2 as min value: ")
		}

    minValue = Min(4, 4)
    if minValue != 4 {
      t.Error("expected 4 as min value: ")
    }

    minValue = Min(-2, -5)
    if minValue != -5 {
      t.Error("expected -5 as min value: ")
    }
}
