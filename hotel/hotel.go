package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var err error
	f, err := os.Open("in.txt")
	if err != nil {
		f = os.Stdin
	}
	r := bufio.NewReader(f)
	var T int
	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		var N, K, M int
		fmt.Fscanln(r, &M, &N, &K)
		var rooms []int
		sum := 0
		for j := 0; j < M; j++ {
			var p int
			fmt.Fscanln(r, &p)
			sum += p
			rooms = append(rooms, p)
		}
		sort.Ints(rooms)
		for j := 0; j < K; j++ {
			sum -= rooms[j]
			sum += N - rooms[j]
		}
		fmt.Println(sum)
	}
}
