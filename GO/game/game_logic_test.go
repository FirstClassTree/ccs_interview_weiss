// in order to run use go test

package game

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateGuess(t *testing.T) {
	/* Example inputs:
	valid: "007", "81", " 10  "
	invalid: "$", "-15", " "
	*/
	guess, err := ValidateGuess("20")
	assert.NoError(t, err)
	assert.Equal(t, 20, guess)
	// "   10 " and such should be legal according to dayna the interviewer
	tests := []struct {
		input    string
		expected int
		err      bool
	}{
		{"20", 20, false},
		{"007", 7, false},
		{"81", 81, false},
		{" 10  ", 10, false},
		{"$", 0, true},
		{"-15", 0, true},
		{" ", 0, true},
		{"101", 0, true},
		{"0", 0, true},
		{"1", 1, false},
		{"100", 100, false},
		{"50", 50, false},
		{"  99", 99, false},
		{"100  ", 100, false},
		{"abc", 0, true},
		{"1.5", 0, true},
		{"-1", 0, true},
		{"  0  ", 0, true},
	}

	for _, test := range tests {
		guess, err := ValidateGuess(test.input)
		if test.err {
			assert.Error(t, err, "input: %s", test.input)
		} else {
			assert.NoError(t, err, "input: %s", test.input)
			assert.Equal(t, test.expected, guess, "input: %s", test.input)
		}
	}

}

func TestValidateGuessCorrectness(t *testing.T) {
	/* In order to test TestValidateGuessCorrectness you must override the 'random generation'
	Use the following static values and validate the result.
	Example inputs:
	input: "007", "81", " 10  "
	invalid: "$", "-15", " "
	*/

	// to override the random generation added support to pick a specific number instead of random genration
	// given specific seed to ensure reproducability

	// Test cases with expected results
	tests := []struct {
		guess            int
		predefinedNumber int
		predefinedPrime  int
		expectedMatch    bool
	}{
		// -1 means irrlevent
		{guess: 6, predefinedNumber: 1, predefinedPrime: 2, expectedMatch: true},  // Exact match
		{guess: 4, predefinedNumber: 2, predefinedPrime: -1, expectedMatch: true}, // Exact match
		{guess: 55, predefinedNumber: 3, predefinedPrime: 15, expectedMatch: false},
	}
	for _, test := range tests {
		isCorrect := ValidateGuessCorrectness(test.guess, &test.predefinedNumber, &test.predefinedPrime)
		assert.Equal(t, test.expectedMatch, isCorrect, "guess: %d, predefinedNumber: %d, predefinedPrime: %d", test.guess, test.predefinedNumber, test.predefinedPrime)
	}
}

// added testss

func TestGeneratePrefix(t *testing.T) {
	tests := []struct {
		guess int
	}{
		{guess: 10},
		{guess: 75},
		{guess: 150},
		{guess: 200},
		{guess: 50},
		{guess: 0},
	}
	for _, test := range tests {
		prefix := GeneratePrefix(test.guess)
		assert.Contains(t, prefix, strconv.Itoa(test.guess), "guess: %d", test.guess)
		if test.guess >= 0 && test.guess <= 50 {
			assert.Contains(t, prefix, "Your guess is within the safe zone!", "guess: %d", test.guess)
		} else if test.guess > 50 && test.guess <= 150 {
			assert.Contains(t, prefix, "Be careful! Your guess is in the uncertain range.", "guess: %d", test.guess)
		} else {
			assert.Contains(t, prefix, "Your guess is in the high-risk zone!", "guess: %d", test.guess)
		}
	}
}
