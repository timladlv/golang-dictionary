package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var words []string
	for input.Scan() {
		in := input.Text()
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
	fmt.Println("done")
}
