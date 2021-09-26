package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, Search([]int{4,5,6,7,0,1,2}, 0), 4)
}

