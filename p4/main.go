package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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

func isValid(k string, f string) bool {
	switch k {
	case "byr":
		i := atoi(f)
		return 1920<=i && i<=2002
	case "iyr":
		i := atoi(f)
		return 2010<=i && i<=2020
	case "eyr":
		i := atoi(f)
		return 2020<=i && i<=2030
	case "hgt":
		pat := regexp.MustCompile(`^(\d+)(cm|in)$`)
		matches := pat.FindAllStringSubmatch(f, -1)
		if len(matches)>0 {
			i := atoi(matches[0][1])
			switch matches[0][2] {
			case "cm":
				return 150<=i && i<=193
			case "in":
				return 59<=i && i<=76
			default:
				return false
			}
		} else {
			return false
		}
	case "hcl":
		pat := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return pat.MatchString(f)
	case "ecl":
		pat := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		return pat.MatchString(f)
	case "pid":
		pat := regexp.MustCompile(`^\d{9}$`)
		return pat.MatchString(f)
	}
	return false
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")

	req := []string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passport := make(map [string] string)

	v1 := 0
	v2 := 0
	for _, s := range split {
		if len(s)==0 {
			required := true
			valid := true
			for _, k := range req {
				if f, h := passport[k]; !h {
					fmt.Printf("*%s: %s ", k, f)
					required = false
					valid = false
					break
				} else {
					valid = valid && isValid(k, f)
					fmt.Printf("%s: %s ", k, f)
				}
			}

			fmt.Printf("%b\n", valid)

			if required {
				v1++
			}

			if valid {
				v2++
			}

			for k := range passport {
				delete(passport, k)
			}

			continue

		}

		items := strings.Split(s, " ")
		for _, x := range items {
			e := strings.Split(x, ":")
			passport[e[0]] = e[1]
		}

	}
	fmt.Println(v1)
	fmt.Println(v2)
}
