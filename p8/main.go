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

type instruction struct {
	op string
	arg int
}

func run(program []instruction) (int, int, bool) {
	seen := make(map[int]bool)
	acc := 0
	s := false
	var line int
	for line = 0; !s; _,s = seen[line] {
		if line==len(program) {
			return line, acc, true
		}
		seen[line] = true
		i := program[line]
		switch i.op {
		case "acc":
			acc += i.arg
			line++
		case "jmp":
			line += i.arg
		case "nop":
			line++
		}
	}

	return line, acc, false
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	program := make([]instruction, len(split)-1)

	for i, s := range split {
		if len(s)==0 {
			continue
		}

		line := strings.Split(s, " ")
		program[i] = instruction{op: line[0], arg: atoi(line[1])}
	}

	_, acc, _ := run(program)

	fmt.Println(acc)

	for i, inst := range program {
		fmt.Printf("%d: ", i)
		switch inst.op {
			case "jmp":
				program[i] = instruction{op: "nop", arg: inst.arg}
			case "nop":
				program[i] = instruction{op: "jmp", arg: inst.arg}
			default:
				fmt.Printf("\n")
				continue
		}

		line, acc, term := run(program)
		if term {
			fmt.Printf("*%d\n",acc)
			break
		} else {
			program[i] = inst
			fmt.Printf("%d %d\n", line, acc)
		}
	}
}
