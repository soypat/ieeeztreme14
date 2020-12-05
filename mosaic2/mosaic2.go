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
	var W, H, A, B, M, C int
	fmt.Fscanln(r, &W, &H, &A, &B, &M, &C)
	if W == 0 || H == 0 {
		fmt.Println(0)
		return
	}
	Nw, Nh := ceilDiv(W, A), ceilDiv(H, B)
	tiles := Nw * Nh
	piles := ceilDiv(tiles, 10)
	var cuth, cutw int
	if H%B != 0 {
		cutw = Nw * A
	}
	if W%A != 0 {
		if cutw != 0 {
			cutw -= (B - H%B)
			cuth = Nh*B - (A - W%A)
		} else {
			cuth = Nh * B
		}
	}
	totalCost := piles*M + (cuth+cutw)*C
	fmt.Println(totalCost)
	// AreaReutilizable := 0
	// if H%B == 0 { // costado puede reutilizarse
	// 	AreaReutilizable += (W % A) * (H / B)
	// }
	// if W%A == 0 { // costado puede reutilizarse
	// 	AreaReutilizable += (H % B) * (W / A)
	// }

	// fmt.Print(total)
}

func ceilDiv(a, b int) int {
	if a%b == 0 {
		return a / b
	} else {
		return a/b + 1
	}
}
