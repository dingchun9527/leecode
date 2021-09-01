package plusTwoNumber

import "testing"

func TestTwoSum(t *testing.T) {
	array := []int{2,7,11,15}
	target := 9
	index := TwoSum(array, target)
	t.Log(index)
}
