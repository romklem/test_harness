package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gonum.org/v1/gonum/mat"
)

// These values will be changed when we find good values (even values only)
const (
	small  = 256
	medium = 512
	large  = 1024
)

func main() {
	rand.Seed(time.Now().UnixNano())

	mx_path := "../testharness/matrices"
	_, er := os.Stat(mx_path)

	if os.IsNotExist(er) {
		err := os.Mkdir(mx_path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	matrices := createMatrices()
	writeMatricesToFile(matrices)

	results := createResults(matrices)
	writeResultsToFile(results)
}

func createMatrices() []mat.Matrix {
	smallValuesVertical := make([]float64, small*small)
	smallValuesHorizontal := make([]float64, small*small)
	smallValuesSquare1 := make([]float64, small*small)
	smallValuesSquare2 := make([]float64, small*small)

	mediumValuesVertical := make([]float64, medium*medium)
	mediumValuesHorizontal := make([]float64, medium*medium)
	mediumValuesSquare1 := make([]float64, medium*medium)
	mediumValuesSquare2 := make([]float64, medium*medium)

	largeValuesVertical := make([]float64, large*large)
	largeValuesHorizontal := make([]float64, large*large)
	largeValuesSquare1 := make([]float64, large*large)
	largeValuesSquare2 := make([]float64, large*large)

	valuesArray := [][]float64{smallValuesVertical, smallValuesHorizontal, smallValuesSquare1, smallValuesSquare2, mediumValuesVertical, mediumValuesHorizontal, mediumValuesSquare1, mediumValuesSquare2, largeValuesVertical, largeValuesHorizontal, largeValuesSquare1, largeValuesSquare2}

	for _, values := range valuesArray {
		for i := range values {
			values[i] = float64(rand.Intn(1000))
		}
	}

	smallMatVertical := mat.NewDense(small*2, small/2, smallValuesVertical)
	smallMatHorizontal := mat.NewDense(small/2, small*2, smallValuesHorizontal)
	smallMatSquare1 := mat.NewDense(small, small, smallValuesSquare1)
	smallMatSquare2 := mat.NewDense(small, small, smallValuesSquare2)

	mediumMatVertical := mat.NewDense(medium*2, medium/2, mediumValuesVertical)
	mediumMatHorizontal := mat.NewDense(medium/2, medium*2, mediumValuesHorizontal)
	mediumMatSquare1 := mat.NewDense(medium, medium, mediumValuesSquare1)
	mediumMatSquare2 := mat.NewDense(medium, medium, mediumValuesSquare2)

	largeMatVertical := mat.NewDense(large*2, large/2, largeValuesVertical)
	largeMatHorizontal := mat.NewDense(large/2, large*2, largeValuesHorizontal)
	largeMatSquare1 := mat.NewDense(large, large, largeValuesSquare1)
	largeMatSquare2 := mat.NewDense(large, large, largeValuesSquare2)

	return []mat.Matrix{smallMatVertical, smallMatHorizontal, smallMatSquare1, smallMatSquare2, mediumMatVertical, mediumMatHorizontal, mediumMatSquare1, mediumMatSquare2, largeMatVertical, largeMatHorizontal, largeMatSquare1, largeMatSquare2}
}

func writeMatricesToFile(matrices []mat.Matrix) {
	// Open file
	f, err := os.Create("../testharness/matrices/matrices.go")
	if err != nil {
		panic(err)
	}

	// Close file when we are done
	defer f.Close()

	f.WriteString("package matrices\n\n")

	// Write matrices to file in format [][]int{...}
	for i, matrix := range matrices {
		f.WriteString(fmt.Sprintf("var M%v = [][]int{", i))

		rows, columns := matrix.Dims()
		for i := 0; i < rows; i++ {
			f.WriteString(fmt.Sprintf("{"))
			for j := 0; j < columns; j++ {
				f.WriteString(fmt.Sprintf("%d", int(matrix.At(i, j))))
				if j != columns-1 {
					f.WriteString(fmt.Sprintf(", "))
				}
			}

			if i < rows-1 {
				f.WriteString(fmt.Sprintf("},"))
			} else {
				f.WriteString(fmt.Sprintf("}}"))
			}
		}

		f.WriteString(fmt.Sprintf("\n"))
		if (i+1)%2 == 0 {
			f.WriteString(fmt.Sprintf("\n"))
		}
	}
	f.WriteString(fmt.Sprintf("\n"))
}

func createResults(matrices []mat.Matrix) []mat.Matrix {
	results := make([]mat.Matrix, 0)

	for i := 0; i < len(matrices); i += 2 {
		n1, _ := matrices[i].Dims()
		_, m2 := matrices[i+1].Dims()
		result := mat.NewDense(n1, m2, nil)
		result.Mul(matrices[i], matrices[i+1])
		results = append(results, result)
	}

	return results
}

func writeResultsToFile(matrices []mat.Matrix) {
	// Open file
	f, err := os.Create("../testharness/matrices/results.go")
	if err != nil {
		panic(err)
	}

	// Close file when we are done
	defer f.Close()

	f.WriteString("package matrices\n\n")

	// Write matrices to file in format [][]int{...}
	for i, matrix := range matrices {
		f.WriteString(fmt.Sprintf("var R%v%v = [][]int{", i*2, i*2+1))

		rows, columns := matrix.Dims()
		for i := 0; i < rows; i++ {
			f.WriteString(fmt.Sprintf("{"))
			for j := 0; j < columns; j++ {
				f.WriteString(fmt.Sprintf("%d", int(matrix.At(i, j))))
				if j != columns-1 {
					f.WriteString(fmt.Sprintf(", "))
				}
			}

			if i < rows-1 {
				f.WriteString(fmt.Sprintf("},"))
			} else {
				f.WriteString(fmt.Sprintf("}}"))
			}
		}

		f.WriteString(fmt.Sprintf("\n"))
	}
}
