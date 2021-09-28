package rotate
// 功能: 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 思路: 实际上是要把第1行变成最后一列, 第二行变成倒数第二列, 以此例推

func Rotate(matrix [][]int)  {

	length := len(matrix)
	array := make([][]int, length)
	for i:=0; i<length; i++ {
		item := make([]int, length)
		array[i] = item
	}

	for i, row := range matrix {
		for j, col := range row {
			array[j][length-i-1] = col
		}
	}
	copy(matrix, array)
}