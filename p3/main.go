package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	i := 3
	v := 0
	rules := [5][2] int {{1,1}, {3,1}, {5,1}, {7,1}, {1,2}}
	pos := [5]int {1, 3, 5, 7, 1}
	count := [5] int {0, 0, 0, 0, 0}

	fmt.Println(split[0])
	for row, s := range split[1:] {
		if len(s)==0 {
			continue
		}

		m := []byte(s)

		if s[i]=='#' {
			m[i] = 'X'
			v++
		} else {
			m[i] = 'O'
		}

		fmt.Println(string(m))
		i = (i+3) % len(split[0])

		for r, x:= range rules {
			right, down := x[0], x[1]
			if (row+1) % down == 0 {
				if s[pos[r]]=='#' {
					count[r]++
				}

				pos[r] = (pos[r] + right) % len(split[0])
			}
		}
	}
	fmt.Println(v)

	res := 1
	for _, c := range count {
		res = res*c
	}

	fmt.Println(res)
}
