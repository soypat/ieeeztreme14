package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		f = os.Stdin
	}
	var N, cb, cp int
	fmt.Fscanln(f, &N, &cb, &cp)
	var Nb, Np int
	for i := 0; i < N; i++ {
		var b, p int
		fmt.Fscanln(f, &b, &p)
		Nb += b
		Np += p
	}
	var total int
	if Nb%10 != 0 {
		total = cb * (Nb/10 + 1)
	} else {
		total = cb * (Nb / 10)
	}
	if Np%10 != 0 {
		total += cp * (Np/10 + 1)
	} else {
		total += cp * (Np / 10)
	}
	fmt.Print(total)
}
