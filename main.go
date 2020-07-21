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
	"os/user"
	"path/filepath"
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

// Words model for the words field
type Words struct {
	Before []string `json:"before"`
	After  []string `json:"after"`
}

func main() {

	wordsFileBytes, err := ReadWordsFile(wordsFilepath)
	if err != nil {
		log.Fatalln(err)
	}

	allTransforms, err := populateWords(wordsFileBytes)
	if err != nil {
		log.Fatalln(err)
	}

	if len(allTransforms) == 0 {
		return
	}

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
	var wordsFilepathDefault string = "/home/%s/.sprinkle/data/words.json"

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	wordsFilepathDefault = fmt.Sprintf(wordsFilepathDefault, currentUser.Username)

	flag.StringVar(&wordsFilepath, "wordsFile", wordsFilepathDefault, "A file with words to transform")

	flag.Parse()
}

func populateWords(wordsFileBytes []byte) (allTransforms []string, err error) {
	var transforms Transforms

	if len(wordsFileBytes) == 0 {
		return
	}

	err = json.Unmarshal(wordsFileBytes, &transforms)
	if err != nil {
		return
	}

	for _, word := range transforms.Words.Before {
		word = word + otherWord
		allTransforms = append(allTransforms, word)
	}

	for _, word := range transforms.Words.After {
		word = otherWord + word
		allTransforms = append(allTransforms, word)
	}
	return
}

// ReadWordsFile Reads the complete words file.
func ReadWordsFile(path string) (fileBytes []byte, err error) {
	exists, err := ExistsFile(path)
	if err != nil {
		return
	}

	if !exists {
		err = os.MkdirAll(filepath.Dir(path), 0777)
		if err != nil {
			return
		}

		_, err = os.Create(path)
		return
	}

	fileBytes, err = ioutil.ReadFile(path)
	return
}

// ExistsFile - Checks if exists a file.
func ExistsFile(path string) (exists bool, err error) {
	if path == "" {
		return
	}

	exists = true

	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = nil
			exists = false
			return
		}
	}

	return
}
