package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []Card
type Card []Side
type Side string

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: flash deckfile\n")
		os.Exit(1)
	}
	d, err := LoadDeck(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading deck: %v\n", err)
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	for {
		c := d[rand.Int()%len(d)]
		ifront := rand.Int() % len(c)
		fmt.Print(c[ifront])
		fmt.Scanln() // wait for user to input newline
		for i, s := range c {
			if i == ifront {
				continue
			}
			fmt.Println(s)
		}
		fmt.Println()
	}
}

func LoadDeck(deckfile string) (Deck, error) {
	f, err := os.Open(deckfile)
	if err != nil {
		return nil, err
	}
	var deck Deck
	s := bufio.NewScanner(f)
	for s.Scan() {
		var line []rune
		var prevc rune
		for _, c := range s.Text() {
			if c == '\t' && prevc == '\t' {
				continue
			}
			line = append(line, c)
			prevc = c
		}
		var card Card
		for _, text := range strings.Split(string(line), "\t") {
			card = append(card, Side(text))
		}
		deck = append(deck, card)
	}
	return deck, nil
}
