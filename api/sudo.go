package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tjuqxy/front-learning/sudo"
)

func SudoCalculate(req *http.Request) (content []byte) {
	param := getPOSTParam(req)
	arrMapI, ok := param["sudo_num"]
	if !ok {
		fmt.Println("no sudo_num")
		return nil
	}

	arrMap, err := Map(arrMapI)
	if err != nil {
		fmt.Println("arr not map, errmsg:", err)
		return nil
	}

	arr := make([][]int, 9)
	for i := 0; i < 9; i++ {
		arr[i] = make([]int, 9)
	}

	for rowStr, vMapI := range arrMap {
		if vMapI == nil {
			fmt.Println(rowStr, "value map is nil")
			continue
		}

		rowInt, err := strconv.Atoi(rowStr)
		if err != nil {
			fmt.Println("convert row to int failed, errmsg:", err)
			continue
		}

		vMap, err := Map(vMapI)
		if err != nil {
			fmt.Println("value map is not map[string]interface{}, errmsg:", err)
			continue
		}

		for colStr, vI := range vMap {
			colInt, err := strconv.Atoi(colStr)
			if err != nil {
				fmt.Println("convert col to int failed, errmsg:", err)
				continue
			}

			vStr, err := String(vI)
			if err != nil {
				fmt.Println("convert value to string failed, errmsg:", err)
				continue
			}

			if len(vStr) != 1 {
				continue
			}

			vInt, err := strconv.Atoi(vStr)
			if err != nil {
				fmt.Println("convert value to int failed, errmsg:", err)
				continue
			}

			arr[rowInt-1][colInt-1] = vInt
		}
	}

	sudo.Cal(arr)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

	ret, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("marshal result to json failed, errmsg:", err)
	}
	return ret
}
