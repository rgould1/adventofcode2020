package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

type Rule map[string]int
type Rules map[string]Rule

func contains(s string, rules *Rules) []string {
	var res []string
	for k, v := range *rules {
		if _, ok := v[s]; ok {
			res = append(res, k)
		}
	}

	return res
}

func mustContain(s string, rules *Rules) int {
	if x, ok := (*rules)[s]; ok {
		res := 1
		for k,v := range x {
			res += v*mustContain(k, rules)
		}
		return res
	} else {
		fmt.Println("can't contain ",s)
		panic("eek")
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	pat := regexp.MustCompile(`(\d+) ([a-z ]+) bags?[,.]`)
	rules := make(Rules)
	for _, s := range split {
		if len(s)==0 {
			continue
		}

		rule := strings.Split(s, " bags contain")
		outer := rule[0]
		matches := pat.FindAllStringSubmatch(rule[1], -1)
		inner := make(Rule)
		for _, m := range matches {
			inner[m[2]] = atoi(m[1])
		}
		rules[outer] = inner
	}

	paths := make(map[string]bool)
	for x := contains("shiny gold", &rules); len(x)!=0; x=x[1:] {
		b := x[0]
		if _, in := paths[b]; !in {
			paths[b] = true
			x = append(x, contains(b, &rules)...)
		}
	}

	fmt.Println(len(paths))

	fmt.Println(mustContain("shiny gold", &rules)-1)

}
