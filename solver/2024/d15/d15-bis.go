package main

import (
	"time"

	"github.com/djlechuck/advent-of-code/utils"
)

type coord struct {
	x, y int
}
type robot struct {
	coord
}

type box struct {
	coord
}

type wall struct {
	coord
}

type grid struct {
	robot      robot
	boxes      []box
	walls      []wall
	maxX, maxY int
}

type moves []string

func main() {
}

func processInput(in utils.Input) (grid, moves) {
	g := grid{}
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
					g.walls = append(g.walls, wall{coord{x, y}})
				case "@":
					g.robot.coord = coord{x, y}
				case "O":
					g.boxes = append(g.boxes, box{coord{x, y}})
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

	g.computeMaxBoundaries()
	for _, move := range m {
		g.robot.move(move, g)
	}

	return 0, time.Since(t).String()
}

func partTwo() (int, string) {
	t := time.Now()

	return 0, time.Since(t).String()
}

func (r *robot) move(dir string, g *grid) {
	nc := r.coord
	switch dir {
	case "^":
		nc.y--
	case "v":
		nc.y++
	case "<":
		nc.x--
	case ">":
		nc.x++
	}
	if g.isWall(nc) {
		return
	}
	b := g.getBoxesInRobotsSight(dir)

	r.coord = nc
}

func (g *grid) getBoxesInRobotsSight(dir string) []*box {
	var boxes []*box
	switch dir {
	case "^":
		for i := g.robot.y; i >= 0; i-- {
			if g.isWall(coord{g.robot.x, i}) {
				break
			}

			if b, ok := g.getBox(coord{g.robot.x, i}); ok {
				boxes = append(boxes, b)
			}
		}
		break
	case "v":
		for i := g.robot.y; i <= g.maxY; i++ {
			if g.isWall(coord{g.robot.x, i}) {
				break
			}

			if b, ok := g.getBox(coord{g.robot.x, i}); ok {
				boxes = append(boxes, b)
			}
		}
		break
	case "<":
		for i := g.robot.x; i >= 0; i-- {
			if g.isWall(coord{i, g.robot.y}) {
				break
			}

			if b, ok := g.getBox(coord{i, g.robot.y}); ok {
				boxes = append(boxes, b)
			}
		}
		break
	case ">":
		for i := g.robot.x; i <= g.maxY; i++ {
			if g.isWall(coord{i, g.robot.y}) {
				break
			}

			if b, ok := g.getBox(coord{i, g.robot.y}); ok {
				boxes = append(boxes, b)
			}
		}
		break
	}

	return boxes
}

func (g *grid) isWall(c coord) bool {
	for _, w := range g.walls {
		if w.coord == c {
			return true
		}
	}
	return false
}

func (g *grid) getBox(c coord) (*box, bool) {
	for _, b := range g.boxes {
		if b.coord == c {
			return &b, true
		}
	}
	return nil, false
}

func (g *grid) computeMaxBoundaries() {
	maxX := 0
	maxY := 0

	for _, w := range g.walls {
		if w.x > maxX {
			maxX = w.x
		}
		if w.y > maxY {
			maxY = w.y
		}
	}

	g.maxX = maxX
	g.maxY = maxY
}

func (g *grid) display() {
	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxY; x++ {
			if g.robot.coord.x == x && g.robot.coord.y == y {
				print("@")
			} else if _, ok := g.getBox(coord{x, y}); ok {
				print("O")
			} else if g.isWall(coord{x, y}) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}
