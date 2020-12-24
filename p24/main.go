package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

type Direction int
const (
	E Direction = iota
	SE
	SW
	W
	NW
	NE
)

// Coords are staggered
//
//        *   (0,2)
//         *  (0,1)
//        *   (0,0)
type Coord struct {
	x, y int
}

func (c Coord) move(d Direction ) Coord {
	ay := c.y
	if ay<0 {
		ay = -ay
	}

	switch d {
	case E:
		return Coord{c.x+1, c.y}
	case W:
		return Coord{c.x-1, c.y}
	case NE:
		return Coord{c.x + ay%2, c.y+1}
	case NW:
		return Coord{c.x - (ay+1)%2, c.y+1}
	case SE:
		return Coord{c.x + ay%2, c.y-1}
	case SW:
		return Coord{c.x - (ay+1)%2, c.y-1}
	}

	return c
}

func (d Direction) str() string {
	switch d {
	case E:
		return "E"
	case W:
		return "W"
	case SW:
		return "SW"
	case SE:
		return "SE"
	case NW:
		return "NW"
	case NE:
		return "NE"
	}

	return ""
}

func (c Coord) adjacent() []Coord {
	res := make([]Coord, 6)
	for i:=E; i<=NE; i++ {
		res[i] = c.move(i)
	}

	return res
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	paths := make([][]Direction,len(split))
	flipped := make(map[Coord]bool)

	for n, s := range split {
		for i:=0; i<len(s); i++ {
			switch s[i] {
			case 'e':
				paths[n] = append(paths[n], E)
			case 's':
				switch s[i+1] {
					case 'e':
						paths[n] = append(paths[n], SE)
					case 'w':
						paths[n] = append(paths[n], SW)
				}
				i++
			case 'w':
				paths[n] = append(paths[n], W)
			case 'n':
				switch s[i+1] {
					case 'e':
						paths[n] = append(paths[n], NE)
					case 'w':
						paths[n] = append(paths[n], NW)
				}
				i++
			}
		}
	}

	for _, p := range paths {
		c := Coord{0, 0}
		if *debug {
			fmt.Printf("(%d, %d) ", c.x, c.y)
		}
		for _, d := range p {
			c = c.move(d)
			if *debug {
				fmt.Printf(" -%s-> (%d, %d)", d.str(), c.x, c.y)
			}
		}

		if flipped[c] {
			if *debug {
				fmt.Printf(" unflip\n")
			}
			delete(flipped, c)
		} else {
			if *debug {
				fmt.Printf(" flip\n")
			}
			flipped[c] = true
		}
	}

	fmt.Println(len(flipped))

	for i:=0; i<100; i++ {
		nf := make(map[Coord]bool)

		for f, _ := range flipped {
			as := f.adjacent()
			b := 0
			for _, a := range as {
				if flipped[a] {
					b++
				}

				bb := 0
				for _, aa := range a.adjacent() {
					if flipped[aa] {
						bb++
					}
				}
				if bb==2 {
					nf[a] = true
				}
			}

			if b==1 || b==2 {
				nf[f] = true
			}
		}

		flipped = nf
	}

	fmt.Println(len(flipped))
}
