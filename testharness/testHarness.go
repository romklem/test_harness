package main

import (
	"fmt"
	"reflect"
	"testharness/matrices"
	"time"
)

type Test struct {
	mat1   [][]int
	mat2   [][]int
	result [][]int
}

func main() {
	tests := []Test{
		{mat1: matrices.M0, mat2: matrices.M1, result: matrices.R01},
		{mat1: matrices.M2, mat2: matrices.M3, result: matrices.R23},
		{mat1: matrices.M4, mat2: matrices.M5, result: matrices.R45},
		{mat1: matrices.M6, mat2: matrices.M7, result: matrices.R67},
		{mat1: matrices.M8, mat2: matrices.M9, result: matrices.R89},
		{mat1: matrices.M10, mat2: matrices.M11, result: matrices.R1011},
	}

	var totalTime int64

	for i, test := range tests {
		var subtestTimeSum int64
		var numIterations int = 10

		testPassed := true

		for subtestNum := 0; subtestNum < numIterations; subtestNum++ {
			start := time.Now()
			result := multiply(test.mat1, test.mat2)
			elapsed := time.Since(start)

			if !reflect.DeepEqual(result, test.result) {
				testPassed = false
				break
			}

			subtestTimeSum += elapsed.Milliseconds()
		}

		fmt.Printf("(%v, %v) x (%v, %v)\n", len(test.mat1), len(test.mat1[0]), len(test.mat2), len(test.mat2[0]))
		if !testPassed {
			fmt.Printf("Error in test %v", i+1)
		} else {
			subtestTimeAvg := subtestTimeSum / int64(numIterations)
			fmt.Printf("Test %v passed in %v ms (average)\n", i+1, subtestTimeAvg)
			totalTime += subtestTimeAvg
		}

		fmt.Println()
	}

	fmt.Printf("\nTotal time: %v ms", totalTime)
}

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	// Your code here

	// The rules are:
	// 1. You can only modify the code of this function (+ other functions you create)
	// 2. You cannot modify the signature of this function
	// 3. You cannot use any external libraries
	// 4. You have to pass all the tests and only these tests

	return nil
}

// func multiply(mat1 [][]int, mat2 [][]int) [][]int {
// 	// Naive implementation
// 	// O(n^3)
// 	result := make([][]int, len(mat1))
// 	for i := range result {
// 		result[i] = make([]int, len(mat2[0]))
// 	}

// 	for i := 0; i < len(mat1); i++ {
// 		for j := 0; j < len(mat2[0]); j++ {
// 			for k := 0; k < len(mat2); k++ {
// 				result[i][j] += mat1[i][k] * mat2[k][j]
// 			}
// 		}
// 	}

// 	return result
// }
