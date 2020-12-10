package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
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

func rtail(ts []string) []string {
	if(len(ts)>0) {
		return ts[:len(ts)-1]
	} else {
		return ts
	}
}

func connections(adapters []int, pos int, m map[int]int) int {
	if n, ok := m[pos]; ok {
		return n
	}

	j := adapters[pos]
	fmt.Println(j)
	var paths int
	for i:=pos+1; i<len(adapters) && adapters[i]<=j+3;i++ {
		if i==len(adapters)-1 {
			paths += 1
		} else {
			paths += connections(adapters, i, m)
		}
	}

	m[pos] = paths

	return paths
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := rtail(strings.Split(contents, "\n"))

	file := make([]int, len(split))

	for i, s := range split {
		file[i] = atoi(s)
	}


	sort.Ints(file)
	sorted := append(file, file[len(file)-1]+3)

	j:=0
	diffs := make(map[int]int)
	for _, a := range sorted {
		diffs[a-j]++
		j = a
	}

	fmt.Println(diffs[1]*diffs[3])

	n := connections(append([]int {0}, file...), 0, make(map[int]int))

	fmt.Println(n)
}
