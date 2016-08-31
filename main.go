package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	d "github.com/timladlv/golang-dictionary/dictionary"
)

func main() {
	fmt.Println("Please enter the dictionary, STOP when you're done")
	input := bufio.NewScanner(os.Stdin)
	var dictionary []string
	for input.Scan() {
		in := strings.ToUpper(input.Text())
		if in == "STOP" {
			break
		} else {
			dictionary = append(dictionary, in)
		}
	}
	alphabet, err := d.CreateAlphabet(dictionary)
	if err != nil {
		fmt.Printf("problem inferring alphabet, error: %q", err.Error())
	}
	fmt.Printf("alphabet is: %q\n", alphabet)
}
