package dictionary

import "errors"

var ErrMissingWords = errors.New("a dictionary must contain at least one word")

func CreateAlphabet(words []string) ([]rune, error) {
	var err error
	var alphabet []rune
	if len(words) == 0 {
		err = ErrMissingWords
	}
	// TODO replace with full impl
	for _, word := range words {
		for _, r := range word {
			alphabet = append(alphabet, r)
		}
	}
	return alphabet, err
}
