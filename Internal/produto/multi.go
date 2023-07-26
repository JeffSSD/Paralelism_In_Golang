package produto

import (
	"Golang_Project/Internal/read"
	"sync"
)

// Função que retorna multiplicação de matriz
func Multiplicar(result *read.Matrix, matrix1, matrix2 read.Matrix, row, col int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(matrix1[0]); i++ {
		(*result)[row][col] += matrix1[row][i] * matrix2[i][col]
	}
}
