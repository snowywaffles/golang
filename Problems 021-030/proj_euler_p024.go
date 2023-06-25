package main

import (
	"fmt"
	"math"
)

func factorial_of(number int) int {
	multiplication_total := 1
	for i := 2; i <= number; i++ {
		multiplication_total *= i
	}
	return multiplication_total
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	fmt.Println("hi")
	fmt.Println("factorial of 4:", factorial_of(1))
	// 10! = 3628800
	// "0123456789"
	// len(digits_slice) - 1
	// xth_lexico_permutation := 4
	// subtract 1 to perm to make work?
	// small_digits_slice := []int{0, 1, 2, 3}

	digits_slice := []int{0, 1, 2, 3}
	var full_permutation string
	lexographical_number := 1
	remainder := (lexographical_number - 1)

	for len(digits_slice) > 0 {
		permutation_index := math.Floor(float64(remainder / factorial_of(len(digits_slice)-1)))
		perm_digit := digits_slice[int(permutation_index)]
		full_permutation += fmt.Sprint(perm_digit)
		remainder -= int(permutation_index) * factorial_of(len(digits_slice)-1)
		digits_slice = removeElement(digits_slice, int(permutation_index))
	}

	fmt.Println("lexicographic permutation of", lexographical_number, "–", full_permutation)

}
