package dictionary

import (
	"errors"
	"unicode/utf8"
)

var ErrMissingWords = errors.New("a dictionary must contain at least one word")
var ErrInsufficientTerms = errors.New("there are insufficient words to determine the alphabet")

func CreateAlphabet(words []string) ([]rune, error) {
	var alphabet []rune
	if len(words) == 0 {
		return alphabet, ErrMissingWords
	}
	if len(words) == 1 {
		return processSingleWord(alphabet, words[0])
	}
	var runePairs [][2]rune
	for i := 0; i < len(words)-1; i++ {
		runePair, inferred := inferRuneSequenceFromTwoWords(words[i], words[i+1])
		if inferred {
			runePairs = append(runePairs, runePair)
		}
	}
	letters := lettersInAlphabet(words)
	for len(runePairs) > 0 {
		r, err := findRuneNeverSecond(runePairs)
		if err != nil {
			return alphabet, err
		}
		alphabet = append(alphabet, r)
		runePairs = removeRedundant(runePairs, r)
	}
	lastRune, error := lastRune(alphabet, letters)
	if error != nil {
		return alphabet, error
	}
	alphabet = append(alphabet, lastRune)
	return alphabet, nil
}

func lastRune(alphabet []rune, letters []rune) (rune, error) {
	if len(alphabet)+1 != len(letters) {
		return 0, ErrInsufficientTerms
	}
	var lastRune rune
	for _, r := range letters {
		found := false
		for _, a := range alphabet {
			if a == r {
				found = true
			}
		}
		if !found {
			lastRune = r
			break
		}
	}
	return lastRune, nil
}

func removeRedundant(runePairs [][2]rune, r rune) [][2]rune {
	var pairsStillNeeded [][2]rune
	for _, rp := range runePairs {
		if rp[0] != r {
			pairsStillNeeded = append(pairsStillNeeded, rp)
		}
	}
	return pairsStillNeeded
}

func processSingleWord(alphabet []rune, word string) ([]rune, error) {
	if utf8.RuneCountInString(word) == 1 && len(word) == 1 {
		r, _ := utf8.DecodeRuneInString(word[:1])
		alphabet = append(alphabet, r)
		return alphabet, nil
	} else {
		return alphabet, ErrInsufficientTerms
	}
}

func lettersInAlphabet(dictionary []string) []rune {
	uniqueLetters := make(map[rune]bool)
	for _, word := range dictionary {
		for _, r := range word {
			if !uniqueLetters[r] {
				uniqueLetters[r] = true
			}
		}
	}
	i := 0
	letters := make([]rune, len(uniqueLetters))
	for letter := range uniqueLetters {
		letters[i] = letter
		i++
	}
	return letters
}

func inferRuneSequenceFromTwoWords(first string, second string) ([2]rune, bool) {
	var runePair [2]rune
	var firstRunes []rune
	for _, r := range first {
		firstRunes = append(firstRunes, r)
	}
	for i, r := range second {
		// when a different rune is found then capture it and return
		if len(firstRunes) < i {
			// nothing in the first word remains
			break
		}
		if firstRunes[i] != r {
			runePair[0] = firstRunes[i]
			runePair[1] = r
			return runePair, true
		}
	}
	return runePair, false
}

func findRuneNeverSecond(runePairs [][2]rune) (rune, error) {
	var err error
	var answer rune
	count := 0
	for _, rp := range runePairs {
		r := rp[0]
		neverSecond := true
		for _, checkRp := range runePairs {
			if r == checkRp[1] {
				neverSecond = false
				break
			}
		}
		if neverSecond && r != answer {
			answer = r
			count++
		}
	}
	if count != 1 {
		// there should be exactly one rune that is never second to be able to infer an alphabet
		err = ErrInsufficientTerms
	}
	return answer, err
}
