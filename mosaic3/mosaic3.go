package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("in1.txt")
	if err != nil {
		f = os.Stdin
	}
	r := bufio.NewReader(f)
	var N, M, R, C int
	fmt.Fscanln(r, &N, &M, &R, &C)
	var mosaic [][]int
	for i := 0; i < R; i++ {
		row, _ := intScanln(C)
		mosaic = append(mosaic, row)
	}

}

func ceilDiv(a, b int) int {
	if a%b == 0 {
		return a / b
	} else {
		return a/b + 1
	}
}

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
