package main

import (
	"flag"
	"fmt"
	"github.com/emersion/go-vcard"
	"io"
	"log"
	"os"
)

type Card struct {
	vcardData *vcard.Card
}

func (c *Card) toCSV() string {
	card := c.vcardData
	return card.PreferredValue(vcard.FieldFormattedName)
}

// jCard https://tools.ietf.org/html/rfc7095
func (*Card) toJSON() {
}

func parseCards(filename string) []Card {
	result := []Card{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := vcard.NewDecoder(f)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		result = append(result, Card{&card})

	}
	return result
}

func main() {
	var filename = flag.String("input", "", "input filename")
	// TODO support array, glob, etc
	flag.Parse()
	if *filename != "" {
		fmt.Println(*filename)
		cards := parseCards(*filename)
		for _, card := range cards {
			fmt.Println(card.toCSV())
		}
	} else {
		// read from stdin
	}
}
