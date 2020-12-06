package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

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
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	v1 := 0
	v2 := 0
	pat := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	for _, s := range split {
		if len(s)==0 {
			continue
		}
		matches := pat.FindAllStringSubmatch(s, -1)
		match := matches[0]
		min := atoi(match[1])
		max := atoi(match[2])
		c := match[3]
		pwd := match[4]
		o := strings.Count(pwd, c)
		if min <= o && o<=max {
			v1++
		}

		if (pwd[min-1]==c[0]) != (pwd[max-1]==c[0]) {
			v2++
		}
	}
	fmt.Println(v1)
	fmt.Println(v2)


}
