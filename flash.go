package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type Deck []Card

type Card struct {
	Hanzi      string
	Pinyin     string
	Definition string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: flash deckfile")
		os.Exit(1)
	}

	deck, err := LoadDeck(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading deck:", err)
		os.Exit(1)
	}

	deck.Review()
}

func LoadDeck(deckfile string) (Deck, error) {
	f, err := os.Open(deckfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var deck Deck
	re := regexp.MustCompile("([^\t\n]+)\t+([^\t\n]+)\t+([^\t\n]+)")
	for _, match := range re.FindAllStringSubmatch(string(buf), -1) {
		if len(match) != 4 {
			return nil, errors.New("bad row")
		}
		card := Card{
			Hanzi:      match[1],
			Pinyin:     match[2],
			Definition: match[3],
		}
		deck = append(deck, card)
	}
	return deck, nil
}

func (deck Deck) Review() {
	for {
		card := deck[rand.Int()%len(deck)]
		if rand.Int()%2 == 0 {
			// Show Hanzi first.
			fmt.Print(card.Hanzi)
			fmt.Scanln() // wait for user to input newline
			fmt.Println(card.Pinyin)
			fmt.Println(card.Definition)
		} else {
			// Show Definition first.
			fmt.Print(card.Definition)
			fmt.Scanln() // wait for user to input newline
			fmt.Println(card.Hanzi)
			fmt.Println(card.Pinyin)
		}
		go say(card.Hanzi)
		fmt.Println()
	}
}

func say(sentence string) {
	cmd := exec.Command("say")
	cmd.Args = []string{"say", sentence}
	cmd.Run()
}
