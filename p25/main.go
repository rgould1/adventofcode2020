package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func transform(value, subject int) int {
	div := 20201227
	value = value * subject
	value = value % div

	return value
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)


	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	cardPK := atoi(split[0])
	doorPK := atoi(split[1])

	cardL, doorL := -1, -1
	v := 1
	for i:=1; cardL==-1 || doorL==-1; i++ {
		v = transform(v, 7)
		if cardL==-1 && v==cardPK {
			cardL = i
		}
		if doorL==-1 && v==doorPK {
			doorL = i
		}
	}

	fmt.Println(cardL, doorL)

	v =1
	for i:=0; i<cardL; i++ {
		v = transform(v, doorPK)
	}
	fmt.Println(v)

	v =1
	for i:=0; i<doorL; i++ {
		v = transform(v, cardPK)
	}
	fmt.Println(v)
}
