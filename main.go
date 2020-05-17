package main

import (
	"fmt"
	"net/url"
	"time"
)

const format = "20060102"

// URLStrategy represents the Azure IP Ranges URL strategy.
type URLStrategy struct {
	t time.Time
}

// DefaultURLStrategy is the default URL strategy.
var DefaultURLStrategy = &URLStrategy{time.Now()}

// getOffset computes the offset given the actual day and the expected day.
func (s *URLStrategy) getOffset() time.Time {
	t := s.t
	if t.Weekday() != time.Monday {
		var days time.Duration
		if t.Weekday() == time.Sunday {
			days = time.Duration(6 - t.Weekday())
		} else {
			days = time.Duration(t.Weekday() - time.Monday)
		}
		timeDelta := time.Hour * 24 * days
		t = t.Add(-timeDelta)
	}
	return t
}

// URL returns the Azure IP Ranges download URL.
func (s *URLStrategy) URL() *url.URL {
	timeDelta := s.getOffset()
	filename := fmt.Sprintf("ServiceTags_Public_%s.json", timeDelta.Format(format))
	rawurl := fmt.Sprintf("https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/%s", filename)
	u, _ := url.Parse(rawurl)
	return u
}

func main() {
	result := DefaultURLStrategy.URL()
	fmt.Println(result)
}
