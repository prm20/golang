package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of "deck"
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}

	cardValues := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	// you can use the underscore _ in the for loop to ignore the index array key
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		// in this for statement, the i is just the index and the card is the new variable that is looping through the receiver
		//for _, card := range d {
		// you can also use the underscore _ & remove i from Println to remove the index from printing in the array
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}


func (d deck) shuffle() deck {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
   
	// rand.Shuffle accepts an integer which represents the number of elements to shuffle and a swap function that swaps the values of i & j
	// note: the swap function mutates d
	rand.Shuffle(len(d), func(i, j int) {
	  d[i], d[j] = d[j], d[i]
	})
   
	// return the shuffled deck
	return d;
}
  
