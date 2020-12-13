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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func occupied(o [][]rune, r, c int) int {
	rows := len(o)
	cols := len(o[0])
	n := 0
	for i := max(0, r-1); i<=min(rows-1, r+1); i++ {
		for j := max(0, c-1); j<= min(cols-1, c+1); j++ {
			if i==r && j==c {
				continue
			} else {
				if o[i][j]=='#' {
					n++
				}
			}
		}
	}

	return n;
}

func occupied2(o [][]rune, r, c int) int {
	rows := len(o)
	cols := len(o[0])
	n := 0
	for _, rd := range []int {-1,0,1} {
		for _, cd := range []int {-1,0,1} {
			if rd==0 &&cd==0 {
				continue
			}
			inner:
			for i,j := r+rd, c+cd; i>=0 && i<rows && j>=0 && j<cols; i,j=i+rd,j+cd {
				switch o[i][j] {
				case '.':
				case 'L':
					break inner
				case '#':
					n++
					break inner
				}
			}
		}
	}

	return n
}

func tick(o, n [][]rune) bool {
	changed := false
	rows := len(o)
	cols := len(o[0])

	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			switch o[r][c] {
			case '.':
			case 'L':
				if occupied(o, r, c)==0 {
					changed = true
					n[r][c] = '#'
				} else {
					n[r][c] = 'L'
				}
			case '#':
				if occupied(o, r, c)>=4 {
					changed = true
					n[r][c] = 'L'
				} else {
					n[r][c] = '#'
				}
			}
		}
	}

	return changed
}

func tick2(o, n [][]rune) bool {
	changed := false
	rows := len(o)
	cols := len(o[0])

	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			switch o[r][c] {
			case '.':
			case 'L':
				if occupied2(o, r, c)==0 {
					changed = true
					n[r][c] = '#'
				} else {
					n[r][c] = 'L'
				}
			case '#':
				if occupied2(o, r, c)>=5 {
					changed = true
					n[r][c] = 'L'
				} else {
					n[r][c] = '#'
				}
			}
		}
	}

	return changed
}


func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	runes := make([][]rune, len(split))
	buf := make([][]rune, len(split))
	for i, s := range split {
		runes[i] = []rune(s)
		buf[i] = []rune(s)
	}


	for ; tick(runes, buf); {
		t := buf
		buf = runes
		runes = t
	}

	occupied := 0
	for _, s  := range buf {
		occupied += strings.Count(string(s), "#")
	}

	fmt.Println(occupied)

	for i, s := range split {
		runes[i] = []rune(s)
		buf[i] = []rune(s)
	}
	for ; tick2(runes, buf); {
		t := buf
		buf = runes
		runes = t
	}

	occupied2 := 0
	for _, s  := range buf {
		occupied2 += strings.Count(string(s), "#")
	}

	fmt.Println(occupied2)

}
