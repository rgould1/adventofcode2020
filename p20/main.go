package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"strconv"
)

var input = flag.String("input", "input.txt", "input file")
var partb = flag.Bool("partb", false, "partb")
var debug = flag.Bool("debug", false, "debug")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type Tile struct {
	image [][]bool
}

func hash(fwd bool, row []bool) int {
	n := len(row)-1
	res := 0
	for i, x := range row {
		if x {
			if fwd {
				res += 1 << uint(n-i)
			} else {
				res += 1 << uint(i)
			}
		}
	}

	return res
}

func (t Tile) top(flip, clockwise bool) int {
	if flip {
		return hash(clockwise, t.image[len(t.image)-1])
	} else {
		return hash(clockwise, t.image[0])
	}
}

func (t Tile) bottom(flip, clockwise bool) int {
	if flip {
		return hash(!clockwise, t.image[0])
	} else {
		return hash(!clockwise, t.image[len(t.image)-1])
	}
}

func (t Tile) left(flip, clockwise bool) int {
	res := make([]bool, len(t.image))
	for i, r := range t.image {
		res[i] = r[0]
	}

	return hash(clockwise==flip, res)
}

func (t Tile) right(flip, clockwise bool) int {
	res := make([]bool, len(t.image))
	for i, r := range t.image {
		res[i] = r[len(t.image)-1]
	}

	return hash(!(clockwise==flip), res)
}

func (t Tile) edge(edge rune, flip, clockwise bool) int {
	switch edge {
	case 't':
		return t.top(flip, clockwise)
	case 'r':
		return t.right(flip, clockwise)
	case 'b':
		return t.bottom(flip, clockwise)
	case 'l':
		return t.left(flip, clockwise)
	}

	return -1
}

type EdgeId struct {
	id int
	flip bool
	edge rune
}

type Orientation struct {
	id int
	flip bool
	top rune
}

func insert(edges map[int][]EdgeId, h int, id EdgeId) {
	if e, ok := edges[h]; ok {
		edges[h] = append(e, id)
	} else {
		edges[h] = []EdgeId{id}
	}
}

func next(edge rune, rev bool) rune {
	if !rev {
		switch edge {
		case 't':
			return 'r'
		case 'r':
			return 'b'
		case 'b':
			return 'l'
		case 'l':
			return 't'
		}
	} else {
		switch edge {
		case 't':
			return 'l'
		case 'l':
			return 'b'
		case 'b':
			return 'r'
		case 'r':
			return 't'
		}
	}

	return 'x'
}

func translate(id Orientation, edge rune) EdgeId {
	var e rune
	switch edge{
	case 't':
		e = id.top
	case 'r':
		e = next(id.top, false)
	case 'b':
		e = next(next(id.top, false), false)
	case 'l':
		e = next(id.top, true)
	}

	return EdgeId{id.id, id.flip, e}
}

func transcoord(top rune, flip bool, width, row, col int) (int, int) {
	var rowp, colp int
	switch top {
	case 't':
		if !flip {
			rowp = row
			colp = col
		} else {
			rowp = width - row
			colp = col
		}
	case 'r':
		if !flip {
			rowp = col
			colp = width - row
		} else {
			rowp = width - col
			colp = width - row
		}
	case 'b':
		if !flip {
			rowp = width - row
			colp = width - col
		} else {
			rowp = row
			colp = width - col
		}
	case 'l':
		if !flip {
			rowp = width - col
			colp = row
		} else {
			rowp = col
			colp = row
		}

	}

	return rowp, colp
}

func corners(ls, rs []EdgeId) []EdgeId {
	res := make([]EdgeId,0)
	for _, a := range ls {
		for _, b := range rs {
			if a.id==b.id && a.flip==b.flip && a.edge==next(b.edge, false) {
				res = append(res, EdgeId{a.id, a.flip, a.edge})
			}
		}
	}

	return res
}

type Coord struct {
	x, y int
}

func contains(grid map[Coord]Orientation, id int) bool {
	for _, x := range grid {
		if x.id==id {
			return true
		}
	}
	return false
}

func Test(pos Coord, grid map[Coord]Orientation, side int, edges map[int][]EdgeId, tiles map[int]Tile) bool {
	if pos.y==0 {
		left := grid[Coord{pos.x-1,0}]
		le := tiles[left.id].edge(translate(left, 'r').edge, left.flip, (pos.x-1)%2==0)

		es := edges[le]

		for _, e := range es {
			if contains(grid, e.id) {
				continue
			}
			grid[pos] = Orientation{e.id, e.flip, next(e.edge, false)}
			if Test(Coord{(pos.x+1) % side, (pos.x+1)/side}, grid, side, edges, tiles) {
				return true
			}
			delete(grid, pos)
		}

	} else if pos.x==0 {
		above := grid[Coord{0, pos.y-1}]
		es := edges[tiles[above.id].edge(translate(above, 'b').edge, above.flip, (pos.y-1)*side%2==0)]
		for _, e := range es {
			if contains(grid, e.id) {
				continue
			}
			grid[pos] = Orientation{e.id, e.flip, e.edge}
			if Test(Coord{1, pos.y}, grid, side, edges, tiles) {
				return true
			}
			delete(grid, pos)
		}
	} else {
		left := grid[Coord{pos.x-1,pos.y}]
		above := grid[Coord{pos.x, pos.y-1}]

		les := edges[tiles[left.id].edge(translate(left, 'r').edge, left.flip, (pos.x-1+pos.y*side)%2==0)]
		aes := edges[tiles[above.id].edge(translate(above, 'b').edge, above.flip, (pos.x+(pos.y-1)*side)%2==0)]

		es := corners(aes, les)

		for _, e := range es {
			if contains(grid, e.id) {
				continue
			}
			grid[pos] = Orientation{e.id, e.flip, e.edge}
			if pos.x == side-1 && pos.y==side-1 || Test(Coord{(pos.x+1) % side, pos.y + (pos.x+1)/side}, grid, side, edges, tiles) {
				return true
			}
			delete(grid, pos)
		}
	}

	return false
}

func monster() (map[Coord]bool, int, int) {
	dat, err := ioutil.ReadFile("monster.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	res := make(map[Coord]bool)
	for r, s := range split {
		for c, x := range s {
			if x=='#' {
				res[Coord{c,r}] = true
			}
		}
	}

	 return res, len(split[0]), len(split)
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	tiles := make(map[int]Tile)

	line := -1
	var id int
	var image [][]bool
	for _, s := range split {
		if s == "" {
			line = -1
			t := Tile{image}
			tiles[id] = t

			if *debug {
				fmt.Printf("%d: Top: %d Bottom: %d Left: %d Right: %d\n", id, t.top(false, true), t.bottom(false, true), t.left(false, true), t.right(false, true))
				fmt.Printf("%d: Top: %d Bottom: %d Left: %d Right: %d\n", id, t.top(false, false), t.bottom(false, false), t.left(false, false), t.right(false, false))
				fmt.Printf("%d: Top: %d Bottom: %d Left: %d Right: %d\n", id, t.top(true, true), t.bottom(true, true), t.left(true, true), t.right(true, true))
				fmt.Printf("%d: Top: %d Bottom: %d Left: %d Right: %d\n", id, t.top(true, false), t.bottom(true, false), t.left(true, false), t.right(true, false))
			}

			continue
		}

		if line == -1 {
			id = atoi(s[5:len(s)-1])
			image = make([][]bool, len(s))
		} else {
			image[line] = make([]bool, len(s))
			for j, c := range s {
				image[line][j] = c=='#'
			}
		}
		line++
	}

	edges := make(map[int][]EdgeId)
	for id, t := range tiles {
		if *debug {
			fmt.Println("tile: ", id)
			for _, r := range t.image {
				for _, c := range r {
					if c {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}
		}

		for _, e := range []rune {'t','l','b','r'} {
			for _, f := range []bool {false, true} {
				for _, d := range []bool {false, true} {

					insert(edges, t.edge(e, f, d), EdgeId{id, f, e})
				}
			}
		}
	}

	side := int(math.Sqrt(float64(len(tiles))))

	grid := make(map[Coord]Orientation)
	outer:
	for id, _ := range tiles {
		grid[Coord{0,0}] = Orientation{id, false, 't'}
		if Test(Coord{1,0}, grid, side, edges, tiles) {
			for y:=0; y<side; y++ {
				for x:=0; x<side; x++ {
					id := grid[Coord{x, y}]
					fmt.Printf("%d %t %c | ", id.id, id.flip, id.top)
				}
				fmt.Printf("\n")
			}
			fmt.Println(grid[Coord{0,0}].id * grid[Coord{side-1,0}].id * grid[Coord{0, side-1}].id * grid[Coord{side-1, side-1}].id)
			break outer
		}
	}

	width := len(tiles[grid[Coord{0,0}].id].image)
	pixels := (width-2)*side
	image = make([][]bool, pixels)
	for y:=0; y<side; y++ {
		for x:=0; x<side; x++ {
			id := grid[Coord{x, y}]
			t := tiles[id.id]
			for row:=1; row<width-1;row++ {
				if x==0 {
					image[y*(width-2) + (row-1)] = make([]bool, pixels)
				}
				for col:=1; col<width-1;col++ {
					rowp, colp := transcoord(id.top, id.flip, width-1, row, col)

					image[y*(width-2) + (row-1)][x*(width-2) + (col-1)] = t.image[rowp][colp]
				}
			}

		}
	}

	if *debug {
		for y:=0; y<pixels; y++ {
			if y % (width-2) == 0 {
				fmt.Printf("\n")
			}
			for x:=0; x<pixels; x++ {
				if x % (width-2) == 0 {
					fmt.Printf(" ")
				}
				if image[y][x] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
	}


	m, mw, mh := monster()
	if *debug {
		for y:=0; y<mh; y++ {
			for x:=0; x<mw; x++ {
				if m[Coord{x,y}] {
					fmt.Printf("O")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}

	maxSightings := 0
	for _, top := range []rune {'t','r','b','l'} {
		for _, flip := range []bool {false, true} {
			sightings := make(map[Coord]bool)
			overlay := make(map[Coord]bool)
			// scan for monsters
			for y:=0; y<pixels-mh; y++ {
				scan:
				for x:=0; x<pixels-mw; x++ {
					mask := make(map[Coord]bool)
					for c, v := range m {
						yp, xp := transcoord(top, flip, pixels-1, y+c.y, x+c.x)
						if image[yp][xp] != v {
							continue scan
						} else {
							mask[Coord{xp, yp}] = true
						}
					}

					// Found
					sightings[Coord{x,y}] = true
					for k,v := range mask {
						overlay[k] = v
					}
				}
			}

			if len(sightings)>maxSightings {
				fmt.Println(len(sightings))
				maxSightings = len(sightings)
				active := 0
				for y:=0; y<pixels; y++ {
					for x:=0; x<pixels; x++ {
						yp, xp := transcoord(top, flip, pixels-1, y, x)
						if image[yp][xp] {
							active++
							if overlay[Coord{xp,yp}] {
								fmt.Printf("O")
							} else {
								fmt.Printf("#")
							}
						} else {
							fmt.Printf(".")
						}
					}
					fmt.Printf("\n")
				}

				fmt.Println(active - len(sightings)*len(m))
			}
		}
	}

}
