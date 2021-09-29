package spiralOrder


func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return []int{}
	}

	result := make([]int, len(matrix)*len(matrix[0]))

	left, right, top, buttom := 0, len(matrix), 0, len(matrix[0])
	for i:=left
	// TODO
}