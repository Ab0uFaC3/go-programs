package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating a type deck of slice of string
type deck []string

// For creating and returning list of playing cards
func newDeck() deck {

	// Declaring an empty deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Making combinations of suits and values
	for _, suit := range cardSuits { // When the var is not used, replace it with '_'
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Printing all the contents of the deck
// 'd' is the actual copy of the elements of type 'deck'
// variable of 'deck' type can access print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Creating a 'hand' of cards
// This takes the deck of cards and number to create hands as input and returns 2 decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converting type 'deck' -> []string -> string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Converting []byte -> string -> []string
func toByte(byteSlice []byte) []string {
	return strings.Split(string(byteSlice), ",")
}

// To save the file to the h/d
// Receiver function is used to call the saveToFile() with a variable of type 'deck'
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

// Read the deck of cards from the h/d
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	// If there is an error while loading the file
	if err != nil {

		// Log the error and exit the program
		fmt.Println("Errror: ", err)
		// Non zero value indicates exit
		os.Exit(1)
	}

	// Converting the stringSlice to deck
	return deck(toByte(byteSlice))

}

// Shuffle the deck of cards
func (d deck) shuffle() {

	// rand shuffles the deck in the same way as it uses same seed value
	//To solve the problem, NewSource() takes current time - time.Now() and converts it into int64 value - UnixNano()
	// This will generate list of different values based on the time
	source := rand.NewSource(time.Now().UnixNano())

	// Once of these values from source is picked by New()
	// New returns a new Rand that uses random values from src to generate other random values.
	random := rand.New(source)

	for i := range d {
		newPosition := random.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
