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
	//split := []string {"FBFBBFFRLR"}

	var m int
	m = 0

	present := make(map [int] bool)

	for _, s := range split {
		if len(s)==0 {
			continue
		}
		row := s[:7]
		col := s[7:]

		r, err := strconv.ParseInt(strings.Replace(strings.Replace(row, "F", "0", -1), "B", "1", -1), 2, 32)
		if err != nil {
			fmt.Printf("Failed to parse: %s, %s, %s\n", s, row, err)
		}

		c, err := strconv.ParseInt(strings.Replace(strings.Replace(col, "L", "0", -1), "R", "1", -1), 2, 32)
		if err != nil {
			fmt.Printf("Failed to parse: %s, %s, %s\n", s, row, err)
		}

		//fmt.Printf("s:%s row:%s col:%s r:%s c:%s\n", s, row, col, r, c)
		id := int(r*8 + c)
		if id > m {
			for i:= m+1; i<id; i++ {
				present[i] = false
			}
			m = id
		}
		present[id] =true
	}

	fmt.Println(m)
	for i:=0; i<m; i++ {
		if !present[i] {
			fmt.Printf("Missing %d\n", i)
		}
	}

}
