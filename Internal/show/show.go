package show

import (
	"Golang_Project/Internal/read"
	"fmt"
	"time"
)

func PrintMatrix1(matrix read.Matrix) {

	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func PrintMatrix(matrix read.Matrix) {

	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Println()
	time.Sleep(time.Second + 180)
}

func init() {
	fmt.Println("====================================")
}
