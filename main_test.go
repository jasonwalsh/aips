package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getOffset(t *testing.T) {
	strategy := &URLStrategy{time.Date(2020, time.May, 16, 0, 0, 0, 0, time.Local)}
	expected := time.Date(2020, time.May, 11, 0, 0, 0, 0, time.Local)
	actual := strategy.getOffset()
	assert.Equal(t, expected, actual)
}
