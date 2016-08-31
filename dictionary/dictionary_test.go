package dictionary

import (
	"reflect"
	"testing"
)

func TestNoWordsReturnsError(t *testing.T) {
	var words []string
	_, uninitialised := CreateAlphabetFromDictionary(words)
	checkMissingWordsError(t, uninitialised)

	_, nilSlice := CreateAlphabetFromDictionary(nil)
	checkMissingWordsError(t, nilSlice)

	words = append(words, "word")
	_, err := CreateAlphabetFromDictionary(words)
	if err != nil {
		t.Error("word passed so should not have errored")
	}
}

func checkMissingWordsError(t *testing.T, err error) {
	if err.Error() != ErrMissingWords.Error() {
		t.Errorf("expected errMissingWords but was %q", err.Error())
	}
}

func TestSingleLetterAlphabet(t *testing.T) {
	expected := []rune{'A'}
	actual, _ := CreateAlphabetFromDictionary([]string{"A"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}
