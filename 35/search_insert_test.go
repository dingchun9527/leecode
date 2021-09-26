package searchInsert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	assert.Equal(t, SearchInsert([]int{1,3,5,6}, 5), 2)
}