package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		in := input.Text()
		if in == "STOP" {
			fmt.Println("enough")
			break
		} else {
			fmt.Println(in)
		}
	}
	fmt.Println("done")
}
