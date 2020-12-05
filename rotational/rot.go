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
	var N, T int
	fmt.Fscanln(r, &N, &T)
	var lights uint64
	for i := 0; i < N; i++ {
		var n uint64
		fmt.Fscan(r, &n)
		fmt.Println(n)
		lights |= 1 << i
	}
	fmt.Println(lights)
}
