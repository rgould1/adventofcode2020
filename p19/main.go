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

type Rule interface {
	valid([]string, map[int]Rule) []string
}

type Char struct {
	c rune
}

type Seq struct {
	s []int
}

type Or struct {
	s1, s2 Seq
}

func (c Char) valid(ss []string, _ map[int]Rule) []string {
	res := make([]string,0)
	for _, s := range ss {
		if len(s)>0 && rune(s[0])==c.c {
			res = append(res, s[1:])
		}
	}

	return res
}

func (seq Seq) valid(ss []string, rules map[int]Rule) []string {
	res := make([]string, 0)
	for _, s := range ss {
		rem := []string{s}
		for _, r := range seq.s {
			rem = rules[r].valid(rem, rules)
		}

		res = append(res, rem...)
	}

	return res
}

func (or Or) valid(ss []string, rules map[int]Rule) []string {
	rem1 := or.s1.valid(ss, rules)
	rem2 := or.s2.valid(ss, rules)

	return append(rem1, rem2...)
}

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

func makeSeq (s string) Seq {
	is := strings.Split(s, " ")
	res := make([]int, len(is))
	for i, s := range is {
		res[i] = atoi(s)
	}

	return Seq{res}
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	i := 0
	rules := make(map[int]Rule)
	for {
		s := split[i]
		if s=="" {
			i++
			break
		}

		parts := strings.Split(s, ": ")
		r := atoi(parts[0])

		if strings.Contains(parts[1], "\"") {
			rules[r] = Char{rune(parts[1][1])}
		} else if strings.Contains(parts[1], "|") {
			or := strings.Split(parts[1], " | ")
			rules[r] = Or{makeSeq(or[0]), makeSeq(or[1])}
		} else {
			rules[r] = makeSeq(parts[1])
		}

		i++
	}

	messages := split[i:]

	valid := 0
	for _, m := range messages {
		for _, rem := range rules[0].valid([]string{m}, rules) {
			if len(rem)==0 {
				valid++
				break
			}
		}
	}

	fmt.Println(valid)
}
