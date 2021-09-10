package removeDuplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 3, RemoveDuplicates([]int{0,1,1,2,2}))
}