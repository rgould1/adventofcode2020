package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
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

func NewSlice(start, end, step int) []int {
    if step <= 0 || end < start {
        return []int{}
    }
    s := make([]int, 0, 1+(end-start)/step)
    for start <= end {
        s = append(s, start)
        start += step
    }
    return s
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n\n")

	pat := regexp.MustCompile(`([a-zA-Z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	reverse := make(map[int](map[string]bool))
	for _, r := range pat.FindAllStringSubmatch(split[0], -1) {
		field := r[1]
		lb := atoi(r[2])
		lt := atoi(r[3])
		rb := atoi(r[4])
		rt := atoi(r[5])

		for _, i := range append(NewSlice(lb, lt, 1), NewSlice(rb, rt, 1)...) {
			if reverse[i] == nil {
				reverse[i] = make(map[string]bool)
			}
			reverse[i][field] = true
		}
	}

	mine := strings.Split(split[1], "\n")
	if mine[0] != "your ticket:" {
		fmt.Println("odd")
		return
	}
	ticket := strings.Split(mine[1], ",")

	nearby := strings.Split(split[2], "\n")
	if nearby[0] != "nearby tickets:" {
		fmt.Println("odd")
		return
	}
	sum := 0
	nearby = nearby[1:len(nearby)-1]
	possible := make([][]string, len(ticket))
	for _, t := range nearby {
		fields := strings.Split(t, ",")
		for i, j := range fields {
			f := atoi(j)
			if fields, ok := reverse[f]; ok {
				newp := make([]string, 0)
				if len(possible[i])==0 {
					for f, _ := range fields {
						newp = append(newp, f)
					}
				} else {
					for _, p := range possible[i] {
						if _, ok := fields[p]; ok {
							newp = append(newp, p)
						}
					}
				}
				possible[i] = newp
			} else {
				sum += f
			}
		}
	}

	positions := make(map[string](map[int]bool))
	for i, fs := range possible {
		for _, f := range fs {
			if positions[f]==nil {
				positions[f] = make(map[int]bool)
			}
			positions[f][i] = true
		}
	}

	fields := make([]string, len(ticket))
	seen := make(map[string]bool)
	for ; len(seen)<len(ticket); {
		for i, fs := range possible {
			if fields[i]!="" {
				continue
			}
			unassigned := make([]string, 0)
			for _, f := range fs {
				if _, ok := seen[f]; !ok {
					unassigned = append(unassigned, f)
				}
			}
			if len(unassigned)==1 {
				fields[i] = unassigned[0]
				seen[unassigned[0]] = true
			}
		}
	}

	fmt.Println(sum)
	product := 1
	for i, ss := range fields {
		if strings.Index(ss, "departure")!=-1 {
			product *= atoi(ticket[i])
		}
	}

	fmt.Println(product)
}
