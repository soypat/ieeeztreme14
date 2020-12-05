// An implementation of Conway's Game of Life.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Next returns the state of the specified cell at the next time step.
func (f *Field) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	if x == 0 {
		alive += b2i(f.Alive(N-1, y))
	} else {
		alive += b2i(f.Alive(x-1, y))
	}
	if y == 0 {
		alive += b2i(f.Alive(x, N-1))
	} else {
		alive += b2i(f.Alive(x, y-1))
	}
	if x == N-1 {
		alive += b2i(f.Alive(0, y))
	} else {
		alive += b2i(f.Alive(x+1, y))
	}
	if y == N-1 {
		alive += b2i(f.Alive(x, 0))
	} else {
		alive += b2i(f.Alive(x, y+1))
	}

	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	// alive == 3 || alive == 2 && f.Alive(x, y)

	return f.Alive(x, y) && (survive[alive] == '1') ||
		!f.Alive(x, y) && (revive[alive] == '1')
	//
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

var revive, survive string

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	// for i := 0; i < (w * h / 4); i++ {
	// 	a.Set(rand.Intn(w), rand.Intn(h), true)
	// }
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	l.a, l.b = l.b, l.a
}

// String returns the game board as a string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte('0')
			if l.a.Alive(x, y) {
				b = '1'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

var N int

func main() {
	f, err := os.Open("in1.txt")
	if err != nil {
		f = os.Stdin
	}
	r := bufio.NewReader(f)
	var rules string
	fmt.Fscanln(r, &rules)
	revive, survive = strings.Split(rules, ";")[0], strings.Split(rules, ";")[1]

	var M int
	fmt.Fscanln(r, &N, &M)
	l := NewLife(N, N)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscanln(r, &s)
		for j, c := range s {
			if c == '1' {
				l.a.Set(j, i, true)
			}
		}
	}
	// fmt.Print("\x0c", l) // Clear screen and print field.
	for i := 0; i < M; i++ {
		l.Step()
		// fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
	fmt.Println(l)
}
