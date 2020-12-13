package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

var input = flag.String("input", "input.txt", "input file")
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

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	earliest := atoi(split[0])
	routes := strings.Split(split[1],",")

	delay := atoi(routes[0])
	route := delay
	ids := make(map[int]int)
	for j, b := range routes {
		if b=="x" {
			continue
		}

		n := atoi(b)

		ids[n] = j
		i := n - earliest % n
		if i<delay {
			delay = i
			route = n
		}
	}

	fmt.Println(delay*route)


        //i[0] * n0 = t
	//i[1] * n1 = t + 1
	//i[2] * n2 = t + 2
	// i[k] % t+k = 0

	minValue := 0
	runningProduct := 1
	for k,v := range ids {
		for (minValue+v)%k != 0 {
			minValue += runningProduct
		}
		runningProduct *= k
		if *debug {
			fmt.Printf("t + %d === 0 mod %d\n", v, k)
			fmt.Printf("Sum so far: %d, product so far: %d\n", minValue, runningProduct)
		}
	}
	fmt.Println(minValue)
}
