package dictionary

import "errors"

var ErrMissingWords = errors.New("a dictionary must contain at least one word")

func CreateAlphabet(words []string) ([]rune, error) {
	var err error
	if len(words) == 0 {
		err = ErrMissingWords
	}
	// TODO replace with impl
	return []rune{'B'}, err
}
