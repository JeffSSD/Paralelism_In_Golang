package main

import (
	"Golang_Project/Internal/produto"
	"Golang_Project/Internal/read"
	"Golang_Project/Internal/show"
	"Golang_Project/Internal/soma"
	"Golang_Project/Internal/subtra"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o número de linhas: ")
	scanner.Scan()
	linha, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Digite o número de colunas: ")
	scanner.Scan()
	coluna, _ := strconv.Atoi(scanner.Text())
	if linha != coluna {
		fmt.Println("As matrizes não podem ser multiplicadas. Dimensões incompatíveis.")
		//return os.Exit(1)
	} else {

		matrixA := read.CreateRandomMatrix(linha, coluna)
		matrixB := read.CreateRandomMatrix(linha, coluna)

		fmt.Println("Matrix A:")
		go show.PrintMatrix1(matrixA)
		fmt.Println("Matrix B:")
		go show.PrintMatrix1(matrixB)

		// Criar matrizes pra guradar resultados de opereções
		productResult := read.CreateRandomMatrix(linha, coluna)
		sumResult := read.CreateRandomMatrix(linha, coluna)
		subtractResult := read.CreateRandomMatrix(linha, coluna)

		var wg sync.WaitGroup

		// Operação em paralelismo com goroutines
		for i := 0; i < linha; i++ {
			for j := 0; j < coluna; j++ {
				wg.Add(3) // Add 3 goroutines per cell for the three operations
				go produto.Multiplicar(&productResult, matrixA, matrixB, i, j, &wg)
				go soma.Somar(&sumResult, matrixA, matrixB, i, j, &wg)
				go subtra.Subtrar(&subtractResult, matrixA, matrixB, i, j, &wg)
			}
		}
		wg.Wait()

		fmt.Println("RESULTADO DE MULTIPLICAÇÃO:")
		show.PrintMatrix(productResult)

		fmt.Println("RESULTADO DA SOMA:")
		show.PrintMatrix(sumResult)

		fmt.Println("RESULTADO DA DIFERENÇA:")
		show.PrintMatrix(subtractResult)
	}
}

//package main
//
//import (
//	"Golang_Project/Internal"
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"sync"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	//Internal.InputMatrixA()
//	//Internal.InputMatrixB()
//	fmt.Print("Digite o número de linhas da matriz A: ")
//	scanner.Scan()
//	linhaA, _ := strconv.Atoi(scanner.Text())
//
//	fmt.Print("Digite o número de colunas da matriz A: ")
//	scanner.Scan()
//	colunaA, _ := strconv.Atoi(scanner.Text())
//
//	matrixA := Internal.ReadMatrix(linhaA, colunaA)
//
//	fmt.Print("Digite o número de linhas da matriz B: ")
//	scanner.Scan()
//	linhaB, _ := strconv.Atoi(scanner.Text())
//
//	fmt.Print("Digite o número de colunas da matriz B: ")
//	scanner.Scan()
//	colunaB, _ := strconv.Atoi(scanner.Text())
//
//	matrixB := Internal.ReadMatrix(linhaB, colunaB)
//
//	fmt.Println("Matriz A:")
//	Internal.PrintMatrix(matrixA)
//
//	fmt.Println("Matriz B:")
//	Internal.PrintMatrix(matrixB)
//
//	result := make(chan [][]int)
//	wg := sync.WaitGroup{}
//	wg.Add(3)
//
//	// Multiplicação de matrizes
//	go Internal.Multiplicar(matrixA, matrixB, result, &wg)
//
//	// Soma de matrizes
//	go Internal.Somar(matrixA, matrixB, result, &wg)
//
//	// Subtração de matrizes
//	go Internal.Subtrar(matrixA, matrixB, result, &wg)
//
//	wg.Wait()
//	close(result)
//
//	//Internal.Show(answer)
//
//	fmt.Println("Resultado da multiplicação de matrizes:")
//	fmt.Println(<-result)
//
//	fmt.Println("Resultado da soma de matrizes:")
//	fmt.Println(<-result)
//
//	fmt.Println("Resultado da subtração de matrizes:")
//	fmt.Println(<-result)
//}
//
////Internal.Show()
//
////func main() {
////	matrix1 := [][]int{{1, 2}, {3, 4}}
////	matrix2 := [][]int{{5, 6}, {7, 8}}
////
////	// Criar canais para receber os resultados das operações
////	multiplicarResultChan := make(chan [][]int)
////	somarResultChan := make(chan [][]int)
////	subtrarResultChan := make(chan [][]int)
////
////	// Usar WaitGroup para aguardar o término de todas as goroutines
////	var wg sync.WaitGroup
////
////	// goroutine para multiplicação
////	wg.Add(1)
////	go Internal.Multiplicar(matrix1, matrix2, multiplicarResultChan, &wg)
////
////	// goroutine para adição
////	wg.Add(2)
////	go Internal.Somar(matrix1, matrix2, somarResultChan, &wg)
////
////	// goroutine para subtração
////	wg.Add(3)
////	go Internal.Subtrar(matrix1, matrix2, subtrarResultChan, &wg)
////
////	// Aguardar o término de todas as goroutines
////	wg.Wait()
////
////	// Obter os resultados das operações
////	multiplicarResultado := <-multiplicarResultChan
////	somarResultado := <-somarResultChan
////	subtrarResultado := <-subtrarResultChan
////
////	// Fechar os canais
////	close(multiplicarResultChan)
////	close(somarResultChan)
////	close(subtrarResultChan)
////
////	// Exibir os resultados
////	fmt.Println("Resultado da multiplicação das matrizes:")
////	fmt.Println(multiplicarResultado)
////
////	fmt.Println("Resultado da soma das matrizes:")
////	fmt.Println(somarResultado)
////
////	fmt.Println("Resultado da subtração das matrizes:")
////	fmt.Println(subtrarResultado)
////}
//
////
////func main() {
////	Internal.Multiplicar(2, 4)
////	Internal.Somar(2, 4)
////	Internal.Subtrar(2, 4)
////	Internal.Transposta(2, 4)
////}
////func main() {
////	// Duas matrizes de exemplo
////	matrix1 := [][]int{{1, 2}, {3, 4}}
////	matrix2 := [][]int{{5, 6}, {7, 8}}
////
////	result := Multip(matrix1, matrix2)
////	if result == nil {
////		fmt.Println("As matrizes não podem ser multiplicadas. Dimensões incompatíveis.")
////	} else {
////		fmt.Println("Resultado da multiplicação das matrizes:")
////		for _, row := range result {
////			fmt.Println(row)
////		}
////	}
////}
