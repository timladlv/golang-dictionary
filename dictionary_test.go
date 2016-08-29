package main

import "testing"

func TestMissingWordsIsError(t *testing.T) {
	var words []string
	_, err := createDictionary(words)
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != errMissingWords.Error() {
		t.Error("expected errMissingWords")
	}
}
