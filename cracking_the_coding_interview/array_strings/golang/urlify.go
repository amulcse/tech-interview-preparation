package main

import "fmt"

func encodeURL(input string) string {

	output := ""

	for _, c := range input {
		cs := string(c)
		if cs == " " {
			output += "%20"
		} else {
			output += cs
		}
	}

	return output

}

func main() {

	input := "Mr John Smith"

	fmt.Println(encodeURL(input))

}
