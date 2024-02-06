package main

import "fmt"

/*Se asume que las matrices estan correctamaente formateadas
y que tienen las dimensiones correctas para propositos del
problema*/

func matrixMultiplication(M [][]int, N [][]int) [][]int {
	if len(M) != len(N[0]) {
		return [][]int{}
	}

	resM := make([][]int, 3, 3)

	for i := 0; i < len(M); i++ {
		resM[i] = make([]int, 3, 3)
		for j := 0; j < len(N[0]); j++ {
			resM[i][j] = 0
			for k := 0; k < len(N); k++ {
				resM[i][j] += M[i][k] * N[k][j]
			}
		}
	}

	return resM
}

// Se asume que la matriz es 3x3 para propositos del problema
func matrixPotentiation(M [][]int, n int) [][]int {
	if n == 0 {
		return [][]int{
			{
				1, 0, 0,
			},
			{
				0, 1, 0,
			},
			{
				0, 0, 1,
			},
		}
	}

	pm := matrixPotentiation(M, n/2)
	if n%2 == 0 {
		return matrixMultiplication(pm, pm)
	} else {
		return matrixMultiplication(M, matrixMultiplication(pm, pm))
	}
}

func perrin(n int) int {

	if n == 0 {
		return 3
	}

	if n == 1 {
		return 0
	}

	if n == 2 {
		return 2
	}

	tm := [][]int{
		{
			0, 1, 1,
		},
		{
			1, 0, 0,
		},
		{
			0, 1, 0,
		},
	}

	tm = matrixPotentiation(tm, n-2)

	return []int{
		3*tm[0][0] + 2*tm[0][1],
		3*tm[1][0] + 2*tm[1][1],
		3*tm[2][0] + 2*tm[2][1],
	}[0]

}

func main() {
	fmt.Println(perrin(0))
}
