package main

import (
	"bytes"
	"crypto/sha256"
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

func hash(p1, p2 Deck) string {
	var buffer bytes.Buffer
	buffer.WriteString("player 1")
	for _, i := range p1 {
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(",")
	}
	buffer.WriteString("player 2")
	for _, i := range p2 {
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(",")
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(buffer.String())))
}

type Deck []int

func (d *Deck) Draw() int {
	c := (*d)[0]
	*d = (*d)[1:]

	return c
}

func (d *Deck) Replace(c ...int) {
	*d = append(*d, c...)
}

func (d Deck) Score() int {
	score := 0
	for i, c := range d {
		score += c * (len(d) - i)
	}

	return score
}

func (d Deck) Cards() int {
	return len(d)
}

func (d Deck) Copy() Deck {
	d2 := make([]int, len(d))

	for i, c := range d {
		d2[i] = c
	}

	return d2
}

func (d Deck) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(d); i++ {
		buffer.WriteString(strconv.Itoa(d[i]))
		if i != len(d)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

func play(d1, d2 *Deck, game int) int {
	if *debug {
		fmt.Printf("=== Game %d ===\n\n", game)
	}

	round := 1
	games := make(map[string]bool)
	for ; d1.Cards() > 0 && d2.Cards() > 0; round++ {
		h := hash(*d1, *d2)
		if *debug {
			fmt.Printf("Hash: %s\n", h)
		}

		if games[h] {
			if *debug {
				fmt.Printf("Repeating decks, player 1 wins!\n\n")
			}
			return 1
		} else {
			games[h] = true
		}

		if *debug {
			fmt.Printf("-- Round %d (Game %d) --\n", round, game)
			fmt.Printf("Player 1's deck: %s\n", d1.String())
			fmt.Printf("Player 2's deck: %s\n", d2.String())
		}

		c1 := d1.Draw()
		c2 := d2.Draw()

		if *debug {
			fmt.Printf("Player 1 plays: %d\n", c1)
			fmt.Printf("Player 2 plays: %d\n", c2)
		}

		if c1 <= d1.Cards() && c2 <= d2.Cards() {
			// Recurse
			s1 := (*d1)[:c1].Copy()
			s2 := (*d2)[:c2].Copy()

			if *debug {
				fmt.Println("Playing a sub-game to determine the winner...\n\n")
			}

			if play(&s1, &s2, game+1) == 1 {
				if *debug {
					fmt.Printf("...anyway, back to game %d.\n Player 1 wins round %d of game %d!\n\n", game, round, game)
				}
				d1.Replace(c1, c2)
			} else {
				if *debug {
					fmt.Printf("...anyway, back to game %d.\n Player 2 wins round %d of game %d!\n\n", game, round, game)
				}
				d2.Replace(c2, c1)
			}

		} else {
			if c1 > c2 {
				d1.Replace(c1, c2)
				if *debug {
					fmt.Printf("Player 1 wins round %d of game %d!\n\n", round, game)
				}
			} else {
				d2.Replace(c2, c1)
				if *debug {
					fmt.Printf("Player 2 wins round %d of game %d!\n\n", round, game)
				}
			}
		}
	}

	if d1.Cards() > 0 {
		if *debug {
			fmt.Printf("The winner of game %d is player 1!\n\n", game)
		}
		return 1
	} else {
		if *debug {
			fmt.Printf("The winner of game %d is player 2!\n\n", game)
		}
		return 2
	}
}

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(*input)
	check(err)

	contents := string(dat)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	p1 := make(Deck, 0)
	p2 := make(Deck, 0)

	player := 1
	for _, s := range split[1:] {
		if s == "" {
			continue
		}

		if s == "Player 2:" {
			player = 2
			continue
		}

		if player == 1 {
			p1 = append(p1, atoi(s))
		} else {
			p2 = append(p2, atoi(s))
		}
	}

	if !*partb {
		for len(p1) > 0 && len(p2) > 0 {
			c1 := p1.Draw()
			c2 := p2.Draw()

			if c1 > c2 {
				p1.Replace(c1, c2)
			} else {
				p2.Replace(c2, c1)
			}
		}

		if len(p1) == 0 {
			fmt.Println(p2.Score())
		} else {
			fmt.Println(p1.Score())
		}
	} else {
		fmt.Println(p1)
		fmt.Println(p2)

		winner := play(&p1, &p2, 1)

		if *debug {
			fmt.Printf("=== Post-game results ===\n")
			fmt.Printf("Player 1's deck: %s\n", p1.String())
			fmt.Printf("Player 2's deck: %s\n", p2.String())
			fmt.Println()
		}

		if winner == 1 {
			fmt.Println(p1.Score())
		} else {
			fmt.Println(p2.Score())
		}
	}
}
