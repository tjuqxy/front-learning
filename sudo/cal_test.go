package sudo

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	arr := make([][]int, 9)
	for i := 0; i < 9; i++ {
		arr[i] = make([]int, 9)
	}

	Cal(arr)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}
