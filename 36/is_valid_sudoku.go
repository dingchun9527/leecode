package isValidSudoku

// 功能: 请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 分析: 通过map来保证数字不重复, 每行扫一次, 每列扫一次, 3X3的矩阵扫一次

func IsValidSudoku(board [][]byte) bool {

	// 扫描行
	for i:=0; i<9; i++ {
		mapping := make(map[byte]struct{})
		for j:=0; j<9; j++ {
			if _, ok := mapping[board[i][j]]; ok {
				return false
			}
			if board[i][j] != '.' {
				mapping[board[i][j]] = struct{}{}
			}
		}
	}

	// 扫描列
	for i:=0; i<9; i++ {
		mapping := make(map[byte]struct{})
		for j:=0; j<9; j++ {
			if _, ok := mapping[board[j][i]]; ok {
				return false
			}
			if board[j][i] != '.' {
				mapping[board[j][i]] = struct{}{}
			}
		}
	}

	// 扫描3X3矩阵
	for i:=0; i<9; i += 3 {
		for j:=0; j<9; j += 3 {
			mapping := make(map[byte]struct{})
			for m:=i; m<i+3; m++ {
				for n:=j; n<j+3; n++ {
					if _, ok := mapping[board[m][n]]; ok {
						return false
					}
					if board[m][n] != '.' {
						mapping[board[m][n]] = struct{}{}
					}
				}
			}
		}
	}
	return true
}

// 方法2, 一次扫完
func IsValidSudokuNew(board [][]byte) bool {
	var rows, columns [9][9]int
	var boxes [3][3][9]int

	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if board[i][j] == '.' {
				continue
			}

			index := board[i][j] - '1'
			rows[i][index]++
			columns[j][index]++
			boxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || boxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}
