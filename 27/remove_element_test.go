package removeElement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	assert.Equal(t, 2, RemoveElement([]int{1,2,2,3,3}, 3))
}