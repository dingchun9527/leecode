package sumOfTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1, l2 *ListNode

	l1 = l1.InsertAtHead(3).InsertAtHead(4).InsertAtHead(2)
	l2 = l2.InsertAtHead(4).InsertAtHead(6).InsertAtHead(5)

	var expect *ListNode
	expect = expect.InsertAtHead(8).InsertAtHead(0).InsertAtHead(7)
	l3 := AddTwoNumbers(l1, l2)
	assert.Equal(t, true, l3.Equal(expect))
}