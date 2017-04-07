package sudo

var (
	isEnd = false
)

func innerCal(curr int, arr [][]int) {
	if isEnd {
		return
	}

	if curr == 81 {
		isEnd = true
		return
	}

	x := curr / 9
	y := curr % 9

	if arr[x][y] != 0 {
		innerCal(curr+1, arr)
		return
	}

	for num := 1; num < 10; num++ {
		if isEnd {
			return
		}

		hasErr := false

		for i := 0; i < 9; i++ {
			if arr[x][i] == num || arr[i][y] == num {
				hasErr = true
				break
			}
		}

		xx := x / 3 * 3
		yy := y / 3 * 3
		for i := 0; i < 3; i++ {
			if hasErr {
				break
			}
			for j := 0; j < 3; j++ {
				if arr[xx+i][yy+j] == num {
					hasErr = true
					break
				}
			}
		}

		if !hasErr {
			arr[x][y] = num
			innerCal(curr+1, arr)
			if isEnd {
				return
			}
			arr[x][y] = 0
		}
	}
}

func Cal(arr [][]int) {
	isEnd = false
	innerCal(0, arr)
}
