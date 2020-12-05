package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("in.txt")
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
		var s string
		fmt.Fscanln(r, &s)
		fmt.Println(solve(s))
	}
}

type Palim struct {
	start, end int
}

func (p Palim) Len() int {
	return p.end - p.start
}

type Palims []*Palim

func (a Palims) String() (s string) {
	for _, v := range a {
		s += fmt.Sprintf("L=%d  s[%d:%d]", v.Len(), v.start, v.end)
		if v != a[len(a)-1] {
			s += "\n"
		}
	}
	return
}
func (a Palims) Len() int      { return len(a) }
func (a Palims) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Palims) Less(i, j int) bool {
	return a[i].end-a[i].start > a[j].end-a[j].start
}

func solve(s string) int {
	if len(s) <= 2 {
		return len(s)
	}

	var palims Palims
	var max1, max2 int
	for i := 1; i < len(s)-1; i++ {
		if i == len(s)-1 {
			break
		}
		if s[i-1] == s[i+1] {
			p := Palim{start: i, end: i + 1}
			for s[p.start-1] == s[p.end] {
				p.start--
				p.end++
				if p.start == 0 || p.end == len(s) {
					break
				}
			}
			if p.Len() >= max1 {
				max1 = p.Len()
				palims = append(palims, &p)
			} else if p.Len() >= max2 {
				max2 = p.Len()
				palims = append(palims, &p)
			}
		}
		if i == len(s)-2 {
			break
		}
		if s[i] == s[i+1] {
			p := Palim{start: i, end: i + 2}
			for s[p.start-1] == s[p.end] {
				p.start--
				p.end++
				if p.start == 0 || p.end == len(s) {
					break
				}
			}
			if p.Len() >= max1 {
				max1 = p.Len()
				palims = append(palims, &p)
			} else if p.Len() >= max2 {
				max2 = p.Len()
				palims = append(palims, &p)
			}
		}
	}
	if len(palims) == 0 {
		return 0
	}
	sort.Sort(palims)
	var L int
	if palims[0].Len() != len(s) {
		L = palims[0].Len() + 1
	} else {
		L = palims[0].Len() - 1
	}
	for i, p1 := range palims {
		if i == len(palims)-1 {
			break
		}
		for _, p2 := range palims[i+1:] {
			l := p1.Len() + p2.Len()
			if l > L && (!(p1.end > p2.start && p1.start < p2.end) ||
				!(p2.start < p1.end && p2.end > p1.start)) {
				L = l
			}
		}
	}

	// fmt.Println(L)
	fmt.Printf("%q\n", s)
	fmt.Println(palims)
	return L
}

/*
3
neevveenxzyyvvyy
odddoipiipi

3
xabcbayabbaz
abcbaabc
abcba
*/
