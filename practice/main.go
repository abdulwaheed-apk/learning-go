package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("You Entered: ", input)

	fmt.Print("What is your age:")

	ageInput, _ := reader.ReadString('\n')

	age, error := strconv.ParseFloat(strings.TrimSpace(ageInput), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Your age is: ", age)
	}

}
