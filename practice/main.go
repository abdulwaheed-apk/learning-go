package main

import (
	"fmt"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Text: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Print("You Entered: ", input)

	// fmt.Print("What is your age:")

	// ageInput, _ := reader.ReadString('\n')

	// age, error := strconv.ParseFloat(strings.TrimSpace(ageInput), 64)

	// if error != nil {
	// 	fmt.Println(error)
	// } else {
	// 	fmt.Println("Your age is: ", age)
	// }

	// i1, i2, i3 := 10, 20, 30
	// intSum := i1 + i2 + i3

	// fmt.Println("Sum is:", intSum)

	// f1, f2, f3 := 20.5, 30.3, 10.01

	// floatSum := f1 + f2 + f3

	// fmt.Println("FLoat sum is: ", floatSum)

	// floatSum = math.Round(floatSum*100) / 100

	// fmt.Println("Accurate sum is: ", floatSum)

	// circleRadius := 15.2
	// circumference := circleRadius * 2 * math.Pi

	// fmt.Printf("Circumference: %.3f\n", circumference)

	n := time.Now()

	parsedTime := n.Format(time.ANSIC)

	fmt.Println("Time is:", n)
	fmt.Println("Parsed time is:", parsedTime)
}
