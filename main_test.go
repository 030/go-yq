package main

import "testing"

func TestScanFile(t *testing.T) {
	keyValue := map[string]string{
		"a":            "abc",
		"b":            "def",
		"c":            "ghi",
		"under_scores": "ensure that keys may contain an underscore",
	}

	for key, value := range keyValue {
		expected := value
		actual := scanFile("test.yaml", key)
		if expected != actual {
			t.Errorf("Value was incorrect, got: %s, want: %s.", actual, expected)
		}
	}
}
