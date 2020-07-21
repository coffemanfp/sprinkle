package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// wordsFilepath is the flag for a different words filepath
var wordsFilepath string

// otherWord is a helper value for to get the value to replace
const otherWord = "*"

// Transforms model for the words file
type Transforms struct {
	Words Words `json:"words"`
}

type Words struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

func main() {
	wordsFileBytes, err := ioutil.ReadFile(wordsFilepath)
	if err != nil {
		log.Fatalln(err)
	}

	var transforms Transforms

	err = json.Unmarshal(wordsFileBytes, &transforms)
	if err != nil {
		log.Fatalln(err)
	}

	var allTransforms []string

	for _, word := range transforms.Words.Before {
		word = word + otherWord
		allTransforms = append(allTransforms, word)
	}

	for _, word := range transforms.Words.After {
		word = otherWord + word
		allTransforms = append(allTransforms, word)
	}

	log.Println(allTransforms)

	rand.Seed(time.Now().UTC().UnixNano())

	// Reading input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		// Select a random transform
		t := allTransforms[rand.Intn(len(allTransforms))]

		// Replace the input and print the result
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

func init() {
	initFlags()
}

func initFlags() {
	flag.StringVar(&wordsFilepath, "wordsFile", "words.json", "A file with words to transform")

	flag.Parse()
}
