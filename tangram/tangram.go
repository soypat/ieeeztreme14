package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
)

func main() {
    f, err := os.Open("ins.txt")
    if err != nil {
        f = os.Stdin
    }
    r := bufio.NewReader(f)

    var T int
    _, err = fmt.Fscanln(r, &T)
    if err != nil {
        panic(err)
    }
    for t := 0; t < T; t++ {
        var N int
        fmt.Fscanln(r, &N)
        solve(N)
    }
}


const charlist = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func solve(N int) {
	type C [62][62]rune
	
	nextStep := func (i,j int) (int, int) {
		switch i {
		case N-1 || 9:
			return i,j+1
		case i:

		}

	}

    b := bytes.Buffer{}

    var c C
	var i, j int
    for ; i!=-1 ; i,j = nextStep(i,j) {

    }

    , err := b.WriteTo(os.Stdout)
    if err != nil {
        panic(err)
    }
}