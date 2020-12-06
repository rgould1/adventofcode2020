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

func or(a, b *map[rune] bool) *map[rune]bool {
	res := make(map[rune]bool)
	if a==nil {
		for k, v := range *b {
			res[k] = v
		}
	} else {
		for k, v := range *a {
			res[k] = v
		}
		for k, v := range *b {
			res[k] = res[k] || v
		}
	}

	return &res
}

func and(a, b *map[rune]bool) *map[rune]bool {
	res := make(map[rune]bool)
	if a==nil {
		for k, v := range *b {
			res[k] = v
		}
	} else {
		for k, v := range *a {
			if x, ok := (*b)[k]; ok {
				res[k] = v && x
			}
		}
		//fmt.Printf("And: %s %s %s\n", pmap(a), pmap(b), pmap(&res))
	}

	return &res
}

func pmap(x *map[rune]bool) string {
	var sb strings.Builder
	for k, v := range *x {
		if v {
			sb.WriteRune(k)
		}
	}

	 return sb.String()
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n\n")

	count1 := 0
	count2 := 0

	for _, g := range split {
		if len(g)==0 {
			continue
		}

		//fmt.Println("********")

		var q1, q2 *map[rune]bool
		q1 = nil
		q2 = nil
		for _, s := range strings.Split(g, "\n") {
			person := make(map[rune]bool)
			//fmt.Printf("uu %s %s %s %s\n", s, pmap(&person), pmap(&q1), pmap(&q2))
			for _, c := range s {
				person[c] = true
			}

			//fmt.Printf("uu %s %s %s %s\n", s, pmap(&person), pmap(&q1), pmap(&q2))
			q1 = or(q1, &person)
			//fmt.Printf("uu %s %s %s %s\n", s, pmap(&person), pmap(&q1), pmap(&q2))
			q2 = and(q2, &person)
			//fmt.Printf("%s %s %s %s\n", s, pmap(&person), pmap(q1), pmap(q2))
		}

		for _, v := range *q1 {
			if v {
				count1++
			}
		}
		for _, v := range *q2 {
			if v {
				count2++
			}
		}
		fmt.Printf("%d %d\n", count1, count2)
	}

	fmt.Println(count1)
	fmt.Println(count2)
}
