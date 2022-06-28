package wordle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordValidation(t *testing.T) {
	var testCases = []struct {
		name string
		guessWord string
		result string
	}{
		{
			name: "test JELLO",
			guessWord: "JELLO",
			result: "Letter O is GREEN",

		},
		{
			name: "test HELJO",
			guessWord: "HELJO",
			result: "Letter J is BLACK",

		},
		{
			name: "test OLLEH",
			guessWord: "OLLEH",
			result: "Letter H is ORANGE",

		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := wordValidation(tc.guessWord)
			assert.Contains(t, result, tc.result)
		})
	}
}