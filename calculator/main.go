package main

// Uncomment this import section to use required Go packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 1st number: ")
	value1, _ := reader.ReadString('\n')

	fmt.Print("Enter 2nd number: ")
	value2, _ := reader.ReadString('\n')

	calculate(value1, value2)

}

func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	// Convert the second string to a float64
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be number")
	}
	// Calculate and return the result

	sum := float1 + float2

	fmt.Printf("Sum is: %.2f\n", sum)
	return sum
}
