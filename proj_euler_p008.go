package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_product(thousand_digit_number string, number_of_consecutive_digits int) int { // outputs: the largest product

	var slice_of_x_numbers []int // initialize an empty slice

	largest_product := 0

	for i := 0; i <= 12; i++ { // 12 is arbitrary

		current_product := 1

		for j := 0; j <= (number_of_consecutive_digits - 1); j++ { // add numbers to the slice

			character, _ := strconv.Atoi(string(thousand_digit_number[i+j]))
			slice_of_x_numbers = append(slice_of_x_numbers, character) // add one number to the slice
			current_product = current_product * character
			fmt.Println(i, j, slice_of_x_numbers, current_product)

		}

		if current_product > largest_product { // update largest product if there is a new highest product
			largest_product = current_product
		}

		slice_of_x_numbers = nil // clear the slice

	}
	return largest_product
}

func main() {

	data, _ := os.ReadFile("1000_digit_number.txt")                               // read the file with the 1000 digit number
	var thousand_digit_number string = strings.ReplaceAll(string(data), "\n", "") // remove newline characters from the data
	fmt.Println(thousand_digit_number)                                            // debugging purposes
	fmt.Println(find_product(thousand_digit_number, 4))

}
