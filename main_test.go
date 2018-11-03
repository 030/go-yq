package main

import "testing"

func TestScanFile(t *testing.T) {
	keyValue := map[string]string{
		"a":                   "abc",
		"b":                   "def",
		"c":                   "ghi",
		"under_scores":        "ensure that keys may contain an underscore",
		"image":               "someimage",
		"MYSQL_ROOT_PASSWORD": "somewordpress",
	}

	for key, value := range keyValue {
		expected := value
		actual := scanFile("test.yaml", key)
		if expected != actual {
			t.Errorf("Value was incorrect 'Check whether het key '%s' resides in the test yaml file', got value: %s, want: %s.", key, actual, expected)
		}
	}
}
