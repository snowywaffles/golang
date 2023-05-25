package main

import "fmt"

// import "math"

func find_multiplier(i int) {

	current_multiplier := i

	for a := (i - 1); a > 0; a-- {

		if current_multiplier%a == 0 {
			current_multiplier = (current_multiplier / a) // if current_multiplier is divisible by a, divide the current multiplier by a
		}

		// fmt.Println("current_multiplier for:", i, "and a:", a, "|", current_multiplier)
	}

	fmt.Println("current_multiplier:", current_multiplier, "| for i:", i)

}

func main() {

	for j := 2; j <= 10; j++ {
		find_multiplier(j)
	}
}
