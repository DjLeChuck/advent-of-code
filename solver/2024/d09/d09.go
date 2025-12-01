package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type blocks []any
type block struct {
	index  int
	id     int
	isFile bool
	size   int
}
type disk []block

func main() {
	in := utils.ParseInput("inputs/2024/09.txt")

	p1v, p1d := partOne(processInput(in))
	fmt.Printf("Part one: %d - elapsed: %s\n", p1v, p1d)
	p2v, p2d := partTwo(processInputPartTwo(in))
	fmt.Printf("Part two: %d - elapsed: %s\n", p2v, p2d)
}

func processInput(in utils.Input) blocks {
	var r blocks
	var id int

	for i, c := range in[0] {
		v := utils.CastInt(string(c))
		if v == 0 {
			continue
		}

		isFile := i%2 == 0

		r = append(r, expand(id, v, isFile)...)

		if isFile {
			id++
		}
	}

	return r
}

func processInputPartTwo(in utils.Input) disk {
	var d disk
	var id int
	var idx int

	for i, c := range in[0] {
		v := utils.CastInt(string(c))
		if v == 0 {
			continue
		}

		isFile := i%2 == 0

		bId := 0
		if isFile {
			bId = id
			id++
		}

		d = append(d, block{
			index:  idx,
			id:     bId,
			isFile: isFile,
			size:   v,
		})

		idx++
	}

	return d
}

func partOne(b blocks) (int, string) {
	t := time.Now()

	for i := len(b) - 1; i >= 0; i-- {
		v := b[i]
		fsi := slices.Index(b, ".")

		if v != "." {
			b[fsi] = v
			b[i] = "."
		}

		if spaceAllEnd(b) {
			break
		}
	}

	s := 0
	for i, v := range b {
		if v == "." {
			break
		}

		s += i * v.(int)
	}

	return s, time.Since(t).String()
}

func partTwo(d disk) (int, string) {
	t := time.Now()

	for i := len(d) - 1; i >= 0; i-- {
		b := d[i]

		if b.isFile {
			//fi, fb := d.firstFreeSpace()
			fi, fb := d.firstFreeSpace(&b)

			fmt.Printf("process %d (size %d)", b.id, b.size)

			if fb == nil {
				fmt.Println(" -> no free space")
				continue
			}

			fmt.Printf(" -> free space %d (space %d)", fb.id, fb.size)
			if b.fitInFirstFreeSpace(fb) {
				fmt.Printf(" -> moved to %d (size %d)", fb.id, fb.size)

				if b.size < fb.size {
					nfb := block{
						index:  0,
						id:     0,
						isFile: false,
						size:   fb.size - b.size,
					}

					d = append(d[:fi], append([]block{b, nfb}, d[fi+1:i]...)...)
					//fb.size -= b.size
					//d[fi] = *fb
				} else {
					d = append(d[:fi], append([]block{b}, d[fi+1:i]...)...)
				}

				for _, nb := range d {
					if !nb.isFile {
						continue
					}
					fmt.Printf(" nb %d (size %d)\n", nb.id, nb.size)
				}

				//d[fi] = b
				//d[i] = *fb
			}
			fmt.Println("")
		}
	}

	s := 0
	for i, b := range d {
		if !b.isFile {
			continue
		}

		s += b.id * b.size * i
	}

	return s, time.Since(t).String()
}

func expand(id, v int, isFile bool) blocks {
	var r blocks

	for i := 0; i < v; i++ {
		if isFile {
			r = append(r, id)
		} else {
			r = append(r, ".")
		}
	}

	return r
}

func spaceAllEnd(b blocks) bool {
	fs := false

	for _, v := range b {
		if v == "." {
			fs = true
		} else if fs {
			return false
		}
	}

	return true
}

func (d disk) firstFreeSpace(b *block) (int, *block) {
	for i, b2 := range d {
		if !b2.isFile && b2.size >= b.size {
			return i, &b2
		}
	}

	return 0, nil
}

func (b *block) fitInFirstFreeSpace(b2 *block) bool {
	if b2.isFile {
		panic("a file can not fit into another file space!")
	}

	return b.size <= b2.size
}
