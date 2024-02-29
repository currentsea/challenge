package main

import (
	"fmt"
	"regexp"
	"strings"
)

func PadNumbers(input string, padding int) string {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(input, -1)
	for i := 0; i < len(numbers); i++ {
		newString := numbers[i]
		if padding > len(numbers[i]) {
			diff := padding - len(numbers[i])
			for x := 0; x < diff; x++ {
				newString = "0" + newString
			}
		}
		input = strings.Replace(input, numbers[i], newString, 1)
	}
	return input
}

func main() {
	fmt.Println(PadNumbers("James Bond 7", 3))
	fmt.Println(PadNumbers("PI=3.14", 2))
	fmt.Println(PadNumbers("It's 3:13pm", 2))
	fmt.Println(PadNumbers("It's 12:13pm", 2))
	fmt.Println(PadNumbers("99UR1337", 6))
}

/*

Take an input of any string and an integer X, return a string left padded with zeros on any whole numbers found in the inputted string to X chars.

Examples:
Inputs (I): "James Bond 7", 3
Output (O): "James Bond 007"

I: "PI=3.14", 2
O: "PI=03.14"

I: "It's 3:13pm", 2
O: "It's 03:13pm"

I: "It's 12:13pm", 2
O: "It's 12:13pm"

I: "99UR1337", 6
O: return "000099UR001337"

*/
