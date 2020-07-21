package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// otherWord is a helper value for to get the value to replace
const otherWord = "*"

// transforms are simple words posibilities
var transforms = []string{
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get " + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Reading input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		// Select a random transform
		t := transforms[rand.Intn(len(transforms))]

		// Replace the input and print the result
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
