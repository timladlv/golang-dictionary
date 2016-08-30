package main

import (
	"reflect"
	"testing"
)

func TestNoWordsReturnsError(t *testing.T) {
	var words []string
	_, uninitialised := createDictionary(words)
	checkMissingWordsError(t, uninitialised)

	_, nilSlice := createDictionary(nil)
	checkMissingWordsError(t, nilSlice)

	words = append(words, "word")
	_, err := createDictionary(words)
	if err != nil {
		t.Error("word passed so should not have errored")
	}
}

func checkMissingWordsError(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != errMissingWords.Error() {
		t.Error("expected errMissingWords")
	}
}

func TestSingleLetterAlphabet(t *testing.T) {
	expected := []rune{'A'}
	actual, _ := createDictionary([]string{"A"})
	if len(expected) != len(actual) {
		t.Errorf("different length of dictionaries, expected %d actual %d", len(expected), len(actual))
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("rune slices not equal expected %q actual %q", expected, actual)
	}
}
