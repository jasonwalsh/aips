package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getOffset(t *testing.T) {
	strategy := &URLStrategy{time.Date(2020, time.May, 17, 0, 0, 0, 0, time.Local)}
	expected := time.Date(2020, time.May, 11, 0, 0, 0, 0, time.Local)
	actual := strategy.getOffset()
	assert.Equal(t, expected, actual)
}

func TestURL(t *testing.T) {
	strategy := &URLStrategy{time.Date(2020, time.May, 16, 0, 0, 0, 0, time.Local)}
	expected := "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20200511.json"
	actual := strategy.URL().String()
	assert.Equal(t, expected, actual)
}
