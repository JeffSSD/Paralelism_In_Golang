package soma

import (
	"Golang_Project/Internal/read"
	"sync"
)

// Função que retorna a soma matriz
func Somar(result *read.Matrix, matrix1, matrix2 read.Matrix, row, col int, wg *sync.WaitGroup) {
	defer wg.Done()
	(*result)[row][col] = matrix1[row][col] + matrix2[row][col]
}
