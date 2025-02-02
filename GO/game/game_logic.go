package game

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ValidateGuess(input string) (int, error) {
	trimmedInput := strings.TrimSpace(input)
	guess, err := strconv.Atoi(trimmedInput)
	if err != nil {
		return 0, errors.New("input is not a valid number")
	}
	if guess < 1 || guess > 100 {
		return 0, errors.New("input must be a number between 1 and 100")
	}
	return guess, err
}

// added capability to not work randomly but define a specific int
func generateCorrectNumber(predefinedNumber *int, predefinedPrime *int) int {
	//generate random number between 1 - 100
	var number int
	if predefinedNumber != nil {
		number = *predefinedNumber
	} else {
		// Generate random number between 1 - 100
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		number = r.Intn(100) + 1
	}

	if number%2 != 0 {
		if predefinedPrime != nil {
			number += *predefinedPrime
		} else {
			primes := []int{2, 3, 5, 7, 11, 13}
			number += primes[rand.Intn(len(primes))]
		}
		// Add a random prime number less than 17

	} else {
		// Reverse the digits of the number i
		reversed := reverseDigits(number)
		number = reversed
	}

	if number >= 100 {
		number /= 2
	} else if number < 50 {
		number *= 2
	}

	return number
}

func reverseDigits(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

// added control over the random number by defining it explicity or nil for random
func ValidateGuessCorrectness(guess int, predefinedNumber *int, predefinedPrime *int) bool {
	//added function to generate the number
	return guess == generateCorrectNumber(predefinedNumber, predefinedPrime)

}

func GeneratePrefix(guess int) string {
	// Initialize a random seed for unpredictable results
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly select one of three different string formats
	formatChoice := r.Intn(3)
	var prefix string

	// Conditional logic based on the guess
	switch formatChoice {
	case 0:
		// Case 0: Format with "selected" or "chosen" depending on the guess's parity (odd/even)
		if guess%2 == 0 {
			prefix = fmt.Sprintf("The number you selected is %d and it is even!", guess)
		} else {
			prefix = fmt.Sprintf("The number you selected is %d and it is odd!", guess)
		}
	case 1:
		// Case 1: Provide a more complex message for numbers greater than 100
		if guess > 100 {
			prefix = fmt.Sprintf("You selected %d, a number greater than 100! Great choice!", guess)
		} else {
			prefix = fmt.Sprintf("You selected %d, which is a small number!", guess)
		}
	case 2:
		// Case 2: Add a random element to the string
		randomFact := rand.Intn(100)
		prefix = fmt.Sprintf("The number %d has a special fact: %d is a random number generated.", guess, randomFact)
	}

	// Add a suffix based on the range of the guess
	if guess >= 0 && guess <= 50 {
		prefix = fmt.Sprintf("%s Your guess is within the safe zone!", prefix)
	} else if guess > 50 && guess <= 150 {
		prefix = fmt.Sprintf("%s Be careful! Your guess is in the uncertain range.", prefix)
	} else {
		prefix = fmt.Sprintf("%s Your guess is in the high-risk zone!", prefix)
	}

	fmt.Sprintf("%s", prefix)
	//added return because it di't return anything
	return prefix
}
