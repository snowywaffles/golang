package main

import (
	"fmt"
	"math/big"
)

func factorial_of_number(number int64) *big.Int {

	// BigInts are required – instead of int64 – because the factorial of
	// numbers > 30 exceeds the integer limit for int64.

	product := big.NewInt(1)
	for i := 1; i <= int(number); i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}

	return product
}

func main() {

	// funnily enough, I came up with this problem myself! I remembered some vague answer / way to use combinatorics
	// and after some pain importing the combin package, I figured out that on an LxL board, running the combination
	// nCr (where) n = 2L and r = L gives the correct answer experimentally. there's probably a better way to do this,
	// but I cannot be bothered. I made a whole Math StackExchange post on it too.

	var L int64 = 20

	final_numerator := factorial_of_number(2 * L) // the numerator is 40!

	denominator1 := factorial_of_number(L)
	denominator2 := factorial_of_number(L)
	final_denominator := big.NewInt(0)
	final_denominator.Mul(denominator1, denominator2)

	valid_paths_on_a_LxL_board := big.NewInt(0)
	valid_paths_on_a_LxL_board.Div(final_numerator, final_denominator)

	fmt.Println(valid_paths_on_a_LxL_board)
}
