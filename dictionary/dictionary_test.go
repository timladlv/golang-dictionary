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
	actual, err := CreateAlphabet([]string{"B", "A"})
	checkErrorAndExpectedVsActual(t, err, expected, actual)
}

func TestTwoLetterAlphabetDifferentSecondLetters(t *testing.T) {
	expected := []rune{'B', 'A'}
	actual, err := CreateAlphabet([]string{"BB", "BA"})
	checkErrorAndExpectedVsActual(t, err, expected, actual)
}

func TestTwoLetterAlphabetDifferentFirstLetters(t *testing.T) {
	expected := []rune{'B', 'A'}
	actual, err := CreateAlphabet([]string{"BA", "A"})
	checkErrorAndExpectedVsActual(t, err, expected, actual)
}

func TestFourLetterAlphabet(t *testing.T) {
	expected := []rune{'Z', 'Y', 'H', 'X', 'T', 'P'}
	actual, err := CreateAlphabet([]string{"ZZXP", "ZYXX", "ZHXY", "HXYY", "XTPP", "TZXT", "TZXP"})
	checkErrorAndExpectedVsActual(t, err, expected, actual)
}

func checkErrorAndExpectedVsActual(t *testing.T, err error, expected []rune, actual []rune) {
	if err != nil {
		t.Errorf("error detected: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}

func TestInsufficientTerms(t *testing.T) {
	_, err := CreateAlphabet([]string{"CB", "CA", "AB"})
	// should be insufficient because B before A, C before A but what about order of B and C?
	if err != ErrInsufficientTerms {
		t.Errorf("expected: ErrInsufficientTerms but was %q", err)
	}
}
