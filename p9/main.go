package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := rtail(strings.Split(contents, "\n"))

	file := make([]int, len(split))

	for i, s := range split {
		file[i] = atoi(s)
	}

	buf := make(map[int]int)
	n := 25
	for _,i := range file[:n] {
		buf[i]++
	}

	var w int
	for i, x := range file[n:] {
		found := false
		for k, n1 := range buf {
			if n1>0 {
				if n2, ok := buf[x-k]; ok {
					if x==2*k && n2>=2 || n2>=1 {
						found = true
						break
					}
				}
			}
		}

		if !found {
			w = x
			break
		}

		buf[file[i]]--
		buf[x]++
	}

	fmt.Println(w)

	s, e := 0, 0
	sum := file[0]
	for ; true; {
		for ; sum<w; {
			e++
			sum+=file[e]
		}

		if sum==w {
			min, max := file[s], file[s]
			for i:=s; i<=e; i++ {
				if file[i] < min {
					min = file[i]
				}
				if file[i] > max {
					max = file[i]
				}
			}
			fmt.Println(min + max)
			break
		}

		sum -= file[s]
		s++
		if e<s {
			e++
			sum += file[e]
		}

		for ; sum>w && e>s; {
			sum-=file[e]
			e--
		}
	}
}
