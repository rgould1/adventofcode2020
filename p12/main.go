package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

var input = flag.String("input", "input.txt", "input file")
var partb = flag.Bool("partb", false, "Part B")

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	facing := 0
	x, y := 0, 0
	wx, wy := 10, 1

	for _, s := range split {
		dir := s[0]
		amp := atoi(s[1:])

		if !*partb {
			switch dir {
			case 'N':
				y += amp
			case 'E':
				x += amp
			case 'S':
				y -= amp
			case 'W':
				x -= amp
			case 'L':
				facing  = (facing + amp) % 360
			case 'R':
				facing = (facing - amp) % 360
				if facing<0 {
					facing = 360 + facing
				}
			case 'F':
				switch facing {
				case 0:
					x += amp
				case 90:
					y += amp
				case 180:
					x -= amp
				case 270:
					y -= amp
				default:
					fmt.Println("Fail: ", amp)
				}
			}
		} else {
			switch dir {
			case 'N':
				wy += amp
			case 'E':
				wx += amp
			case 'S':
				wy -= amp
			case 'W':
				wx -= amp
			case 'L':
				switch amp {
				case 90:
					t := wy
					wy = wx
					wx = -t
				case 180:
					wx = -wx
					wy = -wy
				case 270:
					t := wy
					wy = -wx
					wx = t
				}
			case 'R':
				switch amp {
				case 90:
					t := wy
					wy = -wx
					wx = t
				case 180:
					wx = -wx
					wy = -wy
				case 270:
					t := wy
					wy = wx
					wx = -t
				}
			case 'F':
				x += wx*amp
				y += wy*amp
			}
		}

	}

	fmt.Println(Abs(x) + Abs(y))

}
