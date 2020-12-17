package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

type point struct {
	x, y, z, w int
}

type bounds struct {
	xl, xh, yl, yh, zl, zh int
}

func (p point) neighbours() map[point]bool {
	res := make(map[point]bool)
	var wrange []int
	if !*partb {
		wrange = []int{0}
	} else {
		wrange = []int{-1,0,1}
	}

	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				for _, w := range wrange {
					if x!=0 || y!=0 || z!=0 || w!=0 {
						res[point{p.x+x,p.y+y,p.z+z, p.w+w}] = true
					}
				}
			}
		}
	}

	return res
}

func (b *bounds) grow(p point) {
	if p.x<b.xl {
		b.xl = p.x
	}
	if p.x>b.xh {
		b.xh = p.x
	}
	if p.y<b.yl {
		b.yl = p.y
	}
	if p.y>b.yh {
		b.yh = p.y
	}
	if p.z<b.zl {
		b.zl = p.z
	}
	if p.z>b.zh {
		b.zh = p.z
	}
}

func grid(active map[point]bool, b bounds) {
	for z:=b.zl; z<=b.zh; z++ {
		fmt.Printf("z=%d\n", z)
		for y:=b.yl; y<=b.yh; y++ {
			for x:=b.xl; x<=b.xh; x++ {
				if _, ok := active[point{x,y,z,0}]; ok {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	active := make(map[point]bool)
	var b bounds
	for y, s := range split {
		for x, r := range s {
			if r=='#' {
				p := point{x,y,0,0}
				active[p] = true
				b.grow(p)
			}
		}
	}

	for n:=0; n<6; n++ {
		zone := make(map[point]bool)
		for p, _ := range active {
			zone[p] = true
			b.grow(p)
			for k, v := range p.neighbours() {
			    zone[k] = v
			}
		}

		if *debug {
			fmt.Printf("****** %d\n", n)
			grid(active, b)
		}


		for p, _ := range zone {
			_, isActive := active[p]
			ns :=  p.neighbours()
			for r, _ := range ns {
				if _, ok := active[r]; !ok {
					delete(ns, r)
				}
			}

			if !(len(ns)==3 || isActive && len(ns)==2) {
				delete(zone, p)
			}
		}

		active = zone
	}

	if *debug {
		grid(active, b)
	}

	fmt.Println(len(active))
}
