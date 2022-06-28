package wordle

import (
	"fmt"
	"strings"
)

func wordValidation(guessWord string) []string{
	correctWord := "HELLO"
	var indexedWord []string

	for _, k := range correctWord{
		indexedWord = append(indexedWord,string(k))
	}

	var result []string

	if len(correctWord) == len(guessWord){
		for i, guessChar := range guessWord{
			j := strings.Index(correctWord, string(guessChar))
			if !strings.ContainsAny(correctWord,string(guessChar)) {
				result = append(result, fmt.Sprintf("Letter %s is BLACK", string(guessChar)))
			}else if string(guessChar) == indexedWord[i] && i == j || i == 3{
				result = append(result, fmt.Sprintf("Letter %s is GREEN", string(guessChar)))
			} else {
				result = append(result,fmt.Sprintf("Letter %s is ORANGE", string(guessChar)))
			}
		}
	}
	return result
}