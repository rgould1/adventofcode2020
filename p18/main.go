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

func calc(start int, tokens []string) (int,int) {
	var res int

	i := start
	var op func(a,b int)int
	for ; i<len(tokens); {
		switch tokens[i] {
		case "+":
			op = func(a,b int)int {
				return a+b
			}
		case "*":
			op = func(a,b int)int {
				return a*b
			}
		case "(":
			j, t := calc(i+1, tokens)
			if op==nil {
				res = t
			} else {
				res = op(res, t)
			}
			i=j
			continue
		case ")":
			return i+1, res
		default:
			t := atoi(tokens[i])
			if op==nil {
				res = t
			} else {
				res = op(res, t)
			}
		}

		i++
	}

	return i,res
}

func calcb(start int, tokens []string) (int,int) {
	i := start
	var args []int
	var op func(a,b int)int

	var left int
	for ; i<len(tokens); {
		switch tokens[i] {
		case "+":
			op = func(a,b int)int {
				return a+b
			}
		case "*":
			op = func(a,b int)int {
				args = append(args, a)
				return b
			}
		case "(":
			j, t := calcb(i+1, tokens)
			if op==nil {
				left = t
			} else {
				left = op(left, t)
			}
			i=j
			continue
		case ")":
			for _,x := range args {
				left *= x
			}
			return i+1, left
		default:
			t := atoi(tokens[i])
			if op==nil {
				left = t
			} else {
				left = op(left, t)
			}
		}

		i++
	}

	for _,x := range args {
		left *= x
	}

	return i,left
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	pat := regexp.MustCompile(`(\d+|[+*()])`)
	resA, resB := 0, 0
	for _, s := range split {
		matches := pat.FindAllString(s, -1)

		_, x := calc(0, matches)
		resA += x

		_, y := calcb(0, matches)
		resB += y

		if *debug {
			fmt.Println(s, x, y)
		}
	}

	fmt.Println(resA)
	fmt.Println(resB)
}
