package main

import (
	"fmt"
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type coord struct {
	x, y int
}
type robot struct {
}

type box struct {
}

type wall struct {
}

type object interface {
	isObject()
}

type grid struct {
	objects map[int]map[int]object
	robot   *robot
}

type moves []string

func main() {
	in := utils.ParseInput("inputs/2024/15-2_example.txt")
	g, m := processInput(in)

	p1v, p1d := partOne(&g, m)
	fmt.Printf("Part one: %d - elapsed: %s\n", p1v, p1d)
}

func processInput(in utils.Input) (grid, moves) {
	g := newGrid()
	m := moves{}
	buildGrid := true

	for y, line := range in {
		if "" == line {
			buildGrid = false
			continue
		}

		for x, c := range line {
			ch := string(c)
			if buildGrid {
				switch ch {
				case "#":
					g.addObject(x, y, wall{})
				case "@":
					g.addObject(x, y, robot{})
				case "O":
					g.addObject(x, y, box{})
				case ".":
					g.addFloor(x, y)
				}
			} else {
				m = append(m, ch)
			}
		}
	}

	return g, m
}

func partOne(g *grid, m moves) (int, string) {
	t := time.Now()

	//g.display()

	for _, move := range m {
		g.move(move)
	}

	return 0, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}

//func (r *robot) move(dir string, g *grid) {
//	nc := r.coord
//	switch dir {
//	case "^":
//		nc.y--
//	case "v":
//		nc.y++
//	case "<":
//		nc.x--
//	case ">":
//		nc.x++
//	}
//	if g.isWall(nc) {
//		return
//	}
//	_ = g.getBoxesInRobotsSight(dir)
//
//	r.coord = nc
//}

func newGrid() grid {
	return grid{
		objects: make(map[int]map[int]object),
	}
}

func (g *grid) move(dir string) {
	fmt.Println(g.getRobotAndNeighboringBoxes(dir))
}

func (g *grid) getRobotAndNeighboringBoxes(dir string) []object {
	var o []object
	rc := g.getRobotCoord()
	o = append(o, g.robot)

	switch dir {
	case "^":
		for i := rc.y; i >= 0; i-- {
			obj := g.objects[i][rc.x]
			if _, ok := obj.(box); ok {
				o = append(o, obj)
			}
			if _, ok := obj.(wall); ok {
				break
			}
		}
	case "v":
		for i := rc.y; i < len(g.objects); i++ {
			obj := g.objects[i][rc.x]
			if _, ok := obj.(box); ok {
				o = append(o, obj)
			}
			if _, ok := obj.(wall); ok {
				break
			}
		}
	case "<":
		for i := rc.x; i >= 0; i-- {
			obj := g.objects[rc.y][i]
			if _, ok := obj.(box); ok {
				o = append(o, obj)
			}
			if _, ok := obj.(wall); ok {
				break
			}
		}
	case ">":
		for i := rc.x; i < len(g.objects); i++ {
			obj := g.objects[rc.y][i]
			if _, ok := obj.(box); ok {
				o = append(o, obj)
			}
			if _, ok := obj.(wall); ok {
				break
			}
		}
	}

	return o
}

func (g *grid) addObject(x, y int, o object) {
	if (*g).objects[y] == nil {
		(*g).objects[y] = make(map[int]object)
	}
	(*g).objects[y][x] = o

	if r, ok := o.(robot); ok {
		g.robot = &r
	}
}

func (g *grid) addFloor(x, y int) {
	if (*g).objects[y] == nil {
		(*g).objects[y] = make(map[int]object)
	}
	(*g).objects[y][x] = nil
}

func (g *grid) getRobotCoord() *coord {
	for y, m := range g.objects {
		for x, o := range m {
			if _, ok := o.(robot); ok {
				return &coord{x, y}
			}
		}
	}

	return nil
}

//func (g *grid) getBoxesInRobotsSight(dir string) []*box {
//	var boxes []*box
//	switch dir {
//	case "^":
//		for i := g.robot.y; i >= 0; i-- {
//			if g.isWall(coord{g.robot.x, i}) {
//				break
//			}
//
//			if b, ok := g.getBox(coord{g.robot.x, i}); ok {
//				boxes = append(boxes, b)
//			}
//		}
//		break
//	case "v":
//		for i := g.robot.y; i <= g.maxY; i++ {
//			if g.isWall(coord{g.robot.x, i}) {
//				break
//			}
//
//			if b, ok := g.getBox(coord{g.robot.x, i}); ok {
//				boxes = append(boxes, b)
//			}
//		}
//		break
//	case "<":
//		for i := g.robot.x; i >= 0; i-- {
//			if g.isWall(coord{i, g.robot.y}) {
//				break
//			}
//
//			if b, ok := g.getBox(coord{i, g.robot.y}); ok {
//				boxes = append(boxes, b)
//			}
//		}
//		break
//	case ">":
//		for i := g.robot.x; i <= g.maxY; i++ {
//			if g.isWall(coord{i, g.robot.y}) {
//				break
//			}
//
//			if b, ok := g.getBox(coord{i, g.robot.y}); ok {
//				boxes = append(boxes, b)
//			}
//		}
//		break
//	}
//
//	return boxes
//}

func (w wall) isObject()  {}
func (b box) isObject()   {}
func (r robot) isObject() {}

//func (g *grid) isWall(c coord) bool {
//	for _, w := range g.walls {
//		if w.coord == c {
//			return true
//		}
//	}
//	return false
//}
//
//func (g *grid) getBox(c coord) (*box, bool) {
//	for _, b := range g.boxes {
//		if b.coord == c {
//			return &b, true
//		}
//	}
//	return nil, false
//}
//
//func (g *grid) computeMaxBoundaries() {
//	maxX := 0
//	maxY := 0
//
//	for _, w := range g.walls {
//		if w.x > maxX {
//			maxX = w.x
//		}
//		if w.y > maxY {
//			maxY = w.y
//		}
//	}
//
//	g.maxX = maxX
//	g.maxY = maxY
//}

func (g *grid) display() {
	for x := 0; x < len(g.objects); x++ {
		for y := 0; y < len(g.objects[x]); y++ {
			o := g.objects[x][y]
			switch o.(type) {
			case robot:
				print("@")
			case box:
				print("O")
			case wall:
				print("#")
			default:
				print(".")
			}
		}
		println()
	}
	println()
}
