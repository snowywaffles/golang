package main

func generate_primes(limit int) []int { // returns a slice with prime numbers from 2 - limit.

	var prime_slice []int
	prime_slice = append(prime_slice, 2, 3, 5) // add the first few primes.

	for k := 4; k <= limit; k++ { // for finding prime numbers up to 'limit'.

		is_k_prime := true

		for l := 0; l < len(prime_slice); l++ { // iterate over all of the items in prime_slice.

			if k%prime_slice[l] == 0 {
				is_k_prime = false // if k is divisible by one of the prime numbers in the slice, k is NOT prime.
			}
		}

		if is_k_prime == true {
			prime_slice = append(prime_slice, k) // if k IS prime, add it to prime_slice.

		}
	}

	return prime_slice

}
