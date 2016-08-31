package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	dictionary "github.com/timladlv/golang-dictionary/dictionary"
)

func main() {
	fmt.Println("Please enter words in the dictionary")
	input := bufio.NewScanner(os.Stdin)
	var words []string
	for input.Scan() {
		in := strings.ToUpper(input.Text())
		if in == "STOP" {
			break
		} else {
			words = append(words, in)
		}
	}
	alphabet, err := dictionary.CreateAlphabetFromDictionary(words)
	if err != nil {
		fmt.Printf("problem inferring alphabet, error: %q", err.Error())
	}
	fmt.Printf("alphabet is: %q\n", alphabet)
}
