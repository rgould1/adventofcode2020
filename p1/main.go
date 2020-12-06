package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

	seen := make([]int, len(split))
	for i, s := range split {
		fmt.Println(s)
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse: %s\n", err)
			break
		}
		seen[i] = n
	}

partA:
	for i, n := range seen {
		for j, m := range seen {
			if i >= j {
				continue
			}

			if n+m == 2020 {
				fmt.Println(n * m)
				break partA
			}
		}
	}
partB:
	for i, n := range seen {
		for j, m := range seen {
			for k, o := range seen {

				if i >= j || j >= k {
					continue
				}

				if n+m+o == 2020 {
					fmt.Println(n * m * o)
					break partB
				}
			}
		}
	}

}
