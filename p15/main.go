package main

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
)

var input = flag.String("input", "20,9,11,0,1,2", "input file")
var stop = flag.Int("stop", 2020, "stop")
var debug = flag.Bool("debug", false, "debug")

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	flag.Parse()
	split := strings.Split(*input, ",")

	n := make(map[int]int)
	for i,s := range split[0:len(split)-1] {
		n[atoi(s)] = i+1
		if *debug {
			fmt.Println(i+1, atoi(s))
		}
	}

	last := atoi(split[len(split)-1])
	for i:=len(split); i<*stop; i++ {
		if *debug {
			fmt.Println(i,last)
		}
		var t int
		if j, ok := n[last]; ok {
			t = i - j
		} else {
			t = 0
		}

		n[last] = i
		last = t
	}

	fmt.Println(last)
}
