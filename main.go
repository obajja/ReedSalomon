package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const P = 7  //prime number (
const K = 5  // number of columns (size of the word)
const R = 10 // size of the vector

func main() {
	coding()
}

func coding() {
	nRows := P - 1
	nCols := K

	t := vecSz(nRows, nCols)
	q := t / nCols
	//var matRS [P][K]uint64
	matRS := [P - 1][K]uint64{}
	matrixRS(&matRS)
	fmt.Println(matRS)

	v := make([]uint64, t)
	readVector("vector.txt", v)
	fmt.Println(v)
	vcod := make([]uint64, (P-1)*q)

	for i := R; i < t; i++ {
		v[i] = 0
	}

	t1 := make([]uint64, nCols)
	t2 := make([]uint64, nRows)

	for i := 0; i < q; i++ {
		for j := i * nCols; j < (i+1)*nCols; j++ {
			t1[j-i*nCols] = v[j]
		}
		matrixVectorMultiplication(matRS, t1, t2)
		for j := i * nRows; j < (i+1)*nRows; j++ {
			vcod[j] = t2[j-i*nRows]
		}
	}

	fmt.Println(vcod)
}
func matrixRS(RS *[P - 1][K]uint64) {
	a := primElement()
	for i := 0; i < P-1; i++ {
		for j := 0; j < K; j++ {
			RS[i][j] = power(a, uint64(i)*uint64(j))
		}
	}
	fmt.Println(RS)
}

func primElement() uint64 {
	var a uint64 = 2
	for i := 2; i < P; i++ {
		if multOrder(uint64(i)) == P-1 {
			return uint64(i)
		}
	}
	return a
}

func multOrder(a uint64) uint64 {
	for i := 1; i < P; i++ {
		if power(a, uint64(i)) == 1 {
			return uint64(i)
		}
	}
	return 0
}

func power(a uint64, b uint64) uint64 {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		return power(a*a%P, b/2)
	}
	return a * power(a*a%P, b/2) % P
}

func readVector(fileName string, vector []uint64) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	i := 0
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		vector[i] = uint64(val)
		i++
	}
	modP(vector)
	defer file.Close()
}

func modP(x []uint64) {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] % P
	}
}

func vecSz(nRows int, nCols int) int {
	if nRows%nCols == 0 {
		return nRows
	}
	if nRows%nCols != 0 {
		return nRows + nCols - nRows%nCols
	}
	return 0
}

func matrixVectorMultiplication(A [P - 1][K]uint64, x []uint64, y []uint64) {
	for i := 0; i < P-1; i++ {
		y[i] = 0
		for j := 0; j < K; j++ {
			y[i] += A[i][j] * x[j]
		}
		modP(y)
	}
}
