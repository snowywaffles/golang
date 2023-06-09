package main

import (
	"fmt"
	"math"
)

func is_k_prime(k int) bool { // returns 'true' if the number is prime, and 'false' if the number is not prime.
	if k < 2 {
		return false
	}
	square_root_of_k := math.Ceil(math.Sqrt(float64(k))) // you only have to check up to the square root, which is great for computational time.
	for possible_divisor := 2; possible_divisor <= int(square_root_of_k); possible_divisor++ {
		if k%possible_divisor == 0 {
			return false
		}
	}
	return true
}

func main() {
	current_best_number_of_consecutive_primes := 0
	for a := -999; a <= 999; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			number_of_consecutive_primes := 0
			for never_stop := true; never_stop == true; { // this 'for' loop will only stop once it's broken
				evaluation := (n * n) + (a * n) + b
				if is_k_prime(evaluation) == false {
					if number_of_consecutive_primes > current_best_number_of_consecutive_primes {
						current_best_number_of_consecutive_primes = number_of_consecutive_primes
						fmt.Println("ALERT: new best a:", a, "and b:", b, "with a record of", number_of_consecutive_primes,
							"number of consecutive primes! the product of a and b is:", (a * b))
						break
					} else {
						break
					}
				} else {
					n++
					number_of_consecutive_primes++
				}
			}
		}
	}
}
