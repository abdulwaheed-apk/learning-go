package main

import (
	"bufio"
	"fmt"
	"os"
)

var pageSize = 25

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("You Entered: ", input)
}
