module Golang_Project

go 1.20
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func multiplyMatrices(matrix1 [][]int, matrix2 [][]int, resultChan chan<- [][]int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	rows1, cols1 := len(matrix1), len(matrix1[0])
//	rows2, cols2 := len(matrix2), len(matrix2[0])
//
//	if cols1 != rows2 {
//		// As matrizes não têm dimensões compatíveis para multiplicação
//		resultChan <- nil
//		return
//	}
//
//	result := make([][]int, rows1)
//	for i := 0; i < rows1; i++ {
//		result[i] = make([]int, cols2)
//		for j := 0; j < cols2; j++ {
//			for k := 0; k < cols1; k++ {
//				result[i][j] += matrix1[i][k] * matrix2[k][j]
//			}
//		}
//	}
//
//	resultChan <- result
//}
//
//func addMatrices(matrix1 [][]int, matrix2 [][]int, resultChan chan<- [][]int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	rows1, cols1 := len(matrix1), len(matrix1[0])
//	rows2, cols2 := len(matrix2), len(matrix2[0])
//
//	if rows1 != rows2 || cols1 != cols2 {
//		// As matrizes não têm dimensões compatíveis para soma
//		resultChan <- nil
//		return
//	}
//
//	result := make([][]int, rows1)
//	for i := 0; i < rows1; i++ {
//		result[i] = make([]int, cols1)
//		for j := 0; j < cols1; j++ {
//			result[i][j] = matrix1[i][j] + matrix2[i][j]
//		}
//	}
//
//	resultChan <- result
//}
//
//func subtractMatrices(matrix1 [][]int, matrix2 [][]int, resultChan chan<- [][]int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	rows1, cols1 := len(matrix1), len(matrix1[0])
//	rows2, cols2 := len(matrix2), len(matrix2[0])
//
//	if rows1 != rows2 || cols1 != cols2 {
//		// As matrizes não têm dimensões compatíveis para subtração
//		resultChan <- nil
//		return
//	}
//
//	result := make([][]int, rows1)
//	for i := 0; i < rows1; i++ {
//		result[i] = make([]int, cols1)
//		for j := 0; j < cols1; j++ {
//			result[i][j] = matrix1[i][j] - matrix2[i][j]
//		}
//	}
//
//	resultChan <- result
//}

//func main() {
//	matrix1 := [][]int{{1, 2}, {3, 4}}
//	matrix2 := [][]int{{5, 6}, {7, 8}}
//
//	// Criar canais para receber os resultados das operações
//	multiplyResultChan := make(chan [][]int)
//	addResultChan := make(chan [][]int)
//	subtractResultChan := make(chan [][]int)
//
//	// Usar WaitGroup para aguardar o término de todas as goroutines
//	var wg sync.WaitGroup
//
//	// Iniciar a primeira goroutine para multiplicação
//	wg.Add(1)
//	go multiplyMatrices(matrix1, matrix2, multiplyResultChan, &wg)
//
//	// Iniciar a segunda goroutine para adição
//	wg.Add(1)
//	go addMatrices(matrix1, matrix2, addResultChan, &wg)
//
//	// Iniciar a terceira goroutine para subtração
//	wg.Add(1)
//	go subtractMatrices(matrix1, matrix2, subtractResultChan, &wg)
//
//	// Aguardar o término de todas as goroutines
//	wg.Wait()
//
//	// Obter os resultados das operações
//	multiplyResult := <-multiplyResultChan
//	addResult := <-addResultChan
//	subtractResult := <-subtractResultChan
//
//	// Fechar os canais
//	close(multiplyResultChan)
//	close(addResultChan)
//	close(subtractResultChan)
//
//	// Exibir os resultados
//	fmt.Println("Resultado da multiplicação das matrizes:")
//	fmt.Println(multiplyResult)
//
//	fmt.Println("Resultado da soma das matrizes:")
//	fmt.Println(addResult)
//
//	fmt.Println("Resultado da subtração das matrizes:")
//	fmt.Println(subtractResult)
//}
