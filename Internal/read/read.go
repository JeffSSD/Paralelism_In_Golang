package read

import (
	"math/rand"
	"time"
)

// Definindo uma matriz como uma slice de slices de inteiros
type Matrix [][]int

func randomInt(min, max int) int {
	// Inicializa a semente do gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Gera um número aleatório dentro do intervalo (min, max)
	return min + rand.Intn(max-min+1)
}

func CreateRandomMatrix(linha, coluna int) Matrix {
	//num := randomInt(0, 100)
	matrix := make(Matrix, linha)
	for i := 0; i < linha; i++ {
		matrix[i] = make([]int, coluna)
		for j := 0; j < coluna; j++ {
			num := randomInt(0, 100)
			num -= 1
			matrix[i][j] = num + (i+j)*2
		}
	}
	return matrix
}
