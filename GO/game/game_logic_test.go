package game

import (
	"math/rand"
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
	// to override the random generation we added support to pick a specific rand

	// given specific seed to ensure reproducability
	source := rand.NewSource(1)
	r := rand.New(source)

	// Test cases with expected results
	tests := []struct {
		guess         int
		expectedMatch bool
	}{
		{guess: 10, expectedMatch: false}, // Example case, adjust based on the fixed seed
		{guess: 20, expectedMatch: false}, // Example case, adjust based on the fixed seed
		// Add more test cases as needed
	}

	for _, test := range tests {
		isCorrect := ValidateGuessCorrectness(test.guess, r)
		assert.Equal(t, test.expectedMatch, isCorrect, "guess: %d", test.guess)
	}
}
