package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
