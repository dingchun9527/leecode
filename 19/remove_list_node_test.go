package removeListNode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	assert.Equal(t, (*ListNode)(nil), RemoveNthFromEnd(&ListNode{Val: 1}, 1))
}