package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter some text:\n")
	text, _ := reader.ReadString('\n')

	fmt.Printf("Input was: %v \n", text)
}
