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

func TestSingleWordLengthGreaterOneFails(t *testing.T) {
	_, err := CreateAlphabet([]string{"BA"})
	if err.Error() != ErrInsufficientTerms.Error() {
		t.Errorf("expected ErrInsufficientTerms but was %q", err.Error())
	}
}

func TestTwoLetterAlphabet(t *testing.T) {
	expected := []rune{'B', 'A'}
	actual, _ := CreateAlphabet([]string{"B", "A"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}

func TestTwoLetterAlphabetDifferentFirstLetters(t *testing.T) {
	expected := []rune{'B', 'A'}
	actual, _ := CreateAlphabet([]string{"BA", "A"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}
