package logic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	for testCase := 0; testCase < 100; testCase++ {
		for size := 2; size < 2058; size++ {
			if length := len(RandomString(size)); length != size {
				t.Fatalf("expected random string of size %d, got %d instead", size, length)
			}
		}
	}
}

func TestMakeRandomStringValid(t *testing.T) {
	lengthStr := RandomString(10)
	assert.Equal(t, len(lengthStr), 10)
	validMqID := RandomString(23)
	assert.False(t, strings.Contains(validMqID, "#"))
	assert.False(t, strings.Contains(validMqID, "!"))
	assert.False(t, strings.Contains(validMqID, "\""))
	assert.False(t, strings.Contains(validMqID, "\\"))
	assert.False(t, strings.Contains(validMqID, "+"))
	assert.False(t, strings.Contains(validMqID, "-"))
	assert.False(t, strings.Contains(validMqID, "{"))
	assert.False(t, strings.Contains(validMqID, "}"))
}
