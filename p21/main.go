package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
)

var input = flag.String("input", "input.txt", "input file")
var partb = flag.Bool("partb", false, "partb")
var debug = flag.Bool("debug", false, "debug")

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

type Set map[string]bool

type Food struct {
	ingredients, allergens Set
}

func intersect(a, b Set) Set {
	res := make(Set)
	for ka, _ := range a {
		if _, ok := b[ka]; ok {
			res[ka] = true
		}
	}

	return res
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	foods := make([]Food, len(split))
	for i, s := range split {
		idx := strings.Index(s, " (contains")
		foods[i].ingredients = make(Set)
		foods[i].allergens = make(Set)
		for _, f := range strings.Split(s[:idx], " ") {
			foods[i].ingredients[f] = true
		}
		for _, a := range strings.Split(s[idx+11:len(s)-1], ", ") {
			foods[i].allergens[a] = true
		}

		fmt.Println(foods[i].ingredients, "***", foods[i].allergens)
	}

	allI := make(map[string][]int)
	allA := make(map[string][]int)
	for i, f := range foods {
		for ing, _ := range f.ingredients {
			if l, ok := allI[ing]; ok {
				allI[ing] = append(l, i)
			} else {
				allI[ing] = []int {i}
			}
		}
		for a, _ := range f.allergens {
			if l, ok := allA[a]; ok {
				allA[a] = append(l, i)
			} else {
				allA[a] = []int {i}
			}
		}
	}

	possible := make(map[string]Set)
	for a, fs := range allA {
		is := foods[fs[0]].ingredients
		fmt.Println(a, is)
		for _, i := range fs[1:] {
			is = intersect(is, foods[i].ingredients)
			fmt.Println(a, is)
		}

		possible[a] = is
	}

	allPossible := make(Set)
	for _, fs := range possible {
		for f, _ := range fs {
			allPossible[f] = true
		}
	}

	fmt.Println(allPossible)

	count := 0
	for i, fs := range allI {
		if _, ok := allPossible[i]; !ok {
			count += len(fs)
		}
	}

	fmt.Println(count)

	// correspondence between possible allergen ingredients and possible allergans
	comp := make([][]Set, len(foods))
	for n, f := range foods {
		maybe := make(Set)
		for i, _ := range f.ingredients {
			if _, ok := allPossible[i]; ok {
				maybe[i] = true
			}
		}

		comp[n] = []Set{maybe, f.allergens}
	}

	known := make(map[string]string)

	for ; len(possible)>0; {
		for a, fs := range possible {
			fmt.Println(a, len(fs))
			if len(fs)==1 {
				var f string
				for f = range fs {
				}
				known[a] = f
				delete(possible, a)
				for _, x := range possible {
					delete(x, f)
				}
			}
		}
	}

	allergens := make([]string, len(known))
	i := 0
	for a, _ := range known {
		allergens[i] = a
		i++
	}

	sort.Strings(allergens)

	for _, a := range allergens {
		fmt.Printf("%s,", known[a])
	}
	fmt.Println()

}
