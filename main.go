package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

const twentyFourHours = time.Hour * 24

const format = "20060102"

// URLStrategy represents the Azure IP Ranges URL strategy.
type URLStrategy struct {
	t time.Time
}

// NewURLStrategy returns a new URLStrategy using the current date.
func NewURLStrategy() *URLStrategy {
	return &URLStrategy{time.Now()}
}

// getOffset computes the offset given the actual day and the expected day.
func (s *URLStrategy) getOffset() time.Time {
	t := s.t
	if !(t.Weekday() == time.Monday) {
		timeDelta := time.Duration(t.Weekday()-time.Monday) * twentyFourHours
		t = t.Add(-timeDelta)
	}
	return t
}

// URL returns the Azure IP Ranges download URL.
func (s *URLStrategy) URL() (*url.URL, error) {
	timeDelta := s.getOffset()
	filename := fmt.Sprintf("ServiceTags_Public_%s.json", timeDelta.Format(format))
	rawurl := fmt.Sprintf("https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/%s", filename)
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func main() {
	strategy := NewURLStrategy()
	result, err := strategy.URL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
