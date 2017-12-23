package util

import "testing"

func TestGetMatchable(t *testing.T) {
	result := ToMatchable("Title1 .-~*(special edition)*~-.")
	if result != "title" {
		t.Errorf("Expected result to be 'title', received %s", result)
	}
}

func TestGetMatchableNoEmptyReturns(t *testing.T) {
	result := ToMatchable("\"Title\"")
	if result != "\"title\"" {
		t.Errorf("Expected result to be \"title\", received %s", result)
	}
}
