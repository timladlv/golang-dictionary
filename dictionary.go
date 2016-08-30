package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var errMissingWords = errors.New("a dictionary must contain at least one word")

func main() {
	fmt.Println("Please enter words in the dictionary")
	input := bufio.NewScanner(os.Stdin)
	var words []string
	for input.Scan() {
		in := strings.ToUpper(input.Text())
		if in == "STOP" {
			fmt.Println("words collected are:")
			for _, word := range words {
				fmt.Println(word)
			}
			break
		} else {
			words = append(words, in)
		}
	}
}

func createDictionary(words []string) ([]rune, error) {
	var err error
	if len(words) == 0 {
		err = errMissingWords
	}
	// TODO replace with impl
	return []rune{'B'}, err
}
