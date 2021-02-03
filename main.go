package main

import (
	"flag"
	"fmt"
	"github.com/emersion/go-vcard"
	"io"
	"log"
	"os"
)

// Map of field names to human readable names

// Map of human readable names to fields

type Card struct {
	vcardData *vcard.Card
}

func (c *Card) prettyPrint() string {
	card := *c.vcardData
	for k, v := range card {
		fmt.Print(k, " ")
		for _, n := range v {
			fmt.Println(n.Value)
		}
	}
	fmt.Println("")
	return ""
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
			card.prettyPrint()
		}
	} else {
		// read from stdin
	}
}
