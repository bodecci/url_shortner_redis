package base62

import (
	"errors"
	"math"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func Encode(number uint64) string {
	// define var encodedBuilder of type strings.Builder
	// used here to efficiently build the resulting string.
	var encodedBuilder strings.Builder

	// preallocates space in the builder to optimize memory allocation.
	// It's a performance optimization to reduce the number of memory allocations.
	//in this case,  preallocating space for up to 11 characters
	encodedBuilder.Grow(11)

	//for loop that continues as long as number is greater than zero.
	// In each iteration, the loop divides number by the length of the alphabet (length),
	// and the result is assigned back to number.
	// This loop effectively iterates through each "digit" of the base62 representation.
	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	//declare variable number of type uint64
	var number uint64

	//for loop that iterates over each character in the encoded string. The loop uses two variables,
	//i (the index of the current character) and symbol (the actual character
	for i, symbol := range encoded {
		// In each iteration, this line calculates the position of the
		//current character (symbol) in the alphabet.
		//The function strings.IndexRune returns the index of the first
		//occurrence of the specified rune in the alphabet.
		//If the character is not found, it returns -1.
		alphabeticPostion := strings.IndexRune(alphabet, symbol)

		//checks if the position of the character in the alphabet is -1,
		//indicating that the character is not present in the alphabet
		//If the character is not found in the alphabet, the function returns an error
		//with a message indicating the invalid character and its value as a uint64
		if alphabeticPostion == -1 {
			return uint64(alphabeticPostion), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPostion) * uint64(math.Pow(float64(length), float64(i)))

	}

	return number, nil
}
