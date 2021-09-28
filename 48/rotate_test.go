package rotate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	array := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	Rotate(array)
	assert.Equal(t, 7, array[0][0])
}