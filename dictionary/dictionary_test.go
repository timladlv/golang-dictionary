package dictionary

import (
	"reflect"
	"testing"
)

func TestNoWordsReturnsError(t *testing.T) {
	var words []string
	_, uninitialised := CreateAlphabet(words)
	checkMissingWordsError(t, uninitialised)

	_, nilSlice := CreateAlphabet(nil)
	checkMissingWordsError(t, nilSlice)

	words = append(words, "word")
	_, err := CreateAlphabet(words)
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
	actual, _ := CreateAlphabet([]string{"A"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}
