package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bit uint64

type Bitmask struct {
	bv   Bit
	size int
}

func New(size int, value Bit) (b *Bitmask) {
	b = new(Bitmask)
	b.bv = value
	b.size = size
	return
}

func (b *Bitmask) Add(bit Bit) {
	b.bv |= bit
}

func (b *Bitmask) Remove(bit Bit) {
	b.bv &^= bit
}

func (b *Bitmask) Toggle(bit Bit) {
	b.bv ^= bit
}

func (b *Bitmask) Check(bit Bit) (ok bool) {
	return (b.bv & bit) != 0
}

// func (b *Bitmask) All() {
// 	var t uint64
// 	for i := 0; i < b.size; i++ {
// 		t += uint64(math.Exp2(float64(i)))
// 	}
// 	b.bv |= Bit(t)
// }

// func (b *Bitmask) Empty() {
// 	b.bv = b.bv >> Bit(63)
// }

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	var M int
	fmt.Fscanln(r, &M)

	s := New(21, 0)

	for i := 0; i < M; i++ {
		var function string
		var x Bit
		fmt.Fscan(r, &function)

		if function != "all" && function != "empty" {
			fmt.Fscan(r, &x)
		}

		switch function {
		case "add":
			s.Add(x)
		case "remove":
			s.Remove(x)
		case "check":
			if s.Check(x) {
				fmt.Fprintln(w, "1")
			} else {
				fmt.Fprintln(w, "0")
			}
		case "toggle":
			s.Toggle(x)
		case "all":
			s = New(21, Bit(1<<21-1))
		case "empty":
			s = New(21, Bit(0))
		}

		fmt.Fscanln(r)
	}

}
