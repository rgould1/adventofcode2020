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

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	var mask string
	var floating map[uint]uint
	mem := make(map[int]uint)
	var set, clear uint
	for _, s := range split {
		e := strings.Split(s, " = ")
		if e[0] == "mask" {
			mask = e[1]
			clear = 0
			set = 0
			floating = make(map[uint]uint)
			nx := uint(0)
			for i:=uint(0); i<uint(len(mask)); i++ {
				r := mask[len(mask)-1-int(i)]
				switch r {
				case 'X':
					floating[nx] = i
					nx++
				case '0':
					clear += 1 << i
				case '1':
					set += 1 << i
				}
			}
			if *debug {
				fmt.Printf("Mask %s: %b %b\n", mask, set, clear)
			}
		} else {
			x := atoi(e[0][4:len(e[0])-1])

			y := uint(atoi(e[1]))
			if !*partb {
				mem[x] = (y | set) &^ clear
				if *debug {
					fmt.Printf("mem[%d]: %b -> %b\n", x, y, mem[x])
				}
			} else {
				base := uint(x) | set
				for i:=uint(0); i<1<<uint(len(floating)); i++ {
					addr := base
					for j,n := range floating {
						if i & (1 << j) != 0 {
							addr |= 1 << n
						} else {
							addr &^= 1 << n
						}
					}

					mem[int(addr)] = y

					if *debug {
						fmt.Printf("mem[%d]: %b -> %b\n", addr, x, mem[int(addr)])
					}
				}
			}
		}

	}

	res := uint(0)
	for _, n := range mem {
		res += n
	}

	fmt.Println(res)
}
