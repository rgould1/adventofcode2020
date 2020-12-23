package main

import (
	"flag"
	"fmt"
	"container/list"
	"strconv"
)

var input = flag.String("input", "469217538", "input")
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

func main() {
	flag.Parse()

	//cups := make([]int, len(*input))
	cups := list.New()
	min := atoi(string((*input)[0]))
	max := min
	for _, s := range *input {
		c := atoi(string(s))
		cups.PushBack(c)
		if c>max {
			max = c
		}

		if c<min {
			min = c
		}
	}

	moves := 100
	num := len(*input)
	if *partb {
		moves =10000000

		i := num
		num = 1000000
		for ; i<num; i++ {
			max++
			cups.PushBack(max)
		}
	}

	cache := make(map[int]*list.Element)
	for c:=cups.Front(); c!=nil; c=c.Next() {
		cache[c.Value.(int)] = c
	}

	for m:=0; m<moves; m++ {
		current := cups.Front()
		d1 := current.Next()
		d2 := d1.Next()
		d3 := d2.Next()

		dest := current.Value.(int)-1
		if dest < min {
			dest = max
		}
		for {
			change := false
			c := cups.Front()
			for i:=0; i<3; i++ {
				c = c.Next()
				if dest == c.Value.(int) {
					dest--
					if dest < min {
						dest = max
					}
					change = true
				}
			}

			if !change {
				break
			}
		}

		p := cache[dest]
		cups.MoveAfter(d1, p)
		cups.MoveAfter(d2, d1)
		cups.MoveAfter(d3, d2)

		cups.MoveToBack(current)
	}

	c := cache[1]

	if !*partb {
		for p:=c.Next(); p!=nil; p=p.Next() {
			fmt.Printf("%d", p.Value)
		}
		for p:=cups.Front(); p!=c; p=p.Next() {
			fmt.Printf("%d", p.Value)
		}
		fmt.Printf("\n")
	} else {
		a := c.Next()
		if a==nil {
			a = cups.Front()
		}
		b := a.Next()
		if b==nil {
			b = cups.Front()
		}
		fmt.Println(a.Value.(int)*b.Value.(int))
	}
}
