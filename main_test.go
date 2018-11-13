package main

import (
	"path/filepath"
	"testing"
)

func TestScanFile(t *testing.T) {
	keyValue := map[string]string{
		"a":                   "abc",
		"b":                   "def",
		"c":                   "ghi",
		"under_scores":        "ensureThatKeysMayContainAnUnderscore",
		"image":               "someimage",
		"MYSQL_ROOT_PASSWORD": "somewordpress",
		"thiscontainsadigit1": "helloworld1",
		"alloallo":            "hallohallo",
		"firefox_checksum":    "sha512:49d776",
	}

	for key, value := range keyValue {
		expected := value
		actual := scanFile("test.yaml", key)
		if expected != actual {
			t.Errorf("Value was incorrect 'Check whether the key '%s' resides in the test yaml file', got value: %s, want: %s.", key, actual, expected)
		}
	}
}

func TestDir(t *testing.T) {
	keyValue := map[string]string{
		"path/to/some.yaml": filepath.Join("path", "to"),
		"hello.yaml":        ".",
	}
	for key, value := range keyValue {
		expected := value
		actual := dir(key)
		if expected != actual {
			t.Errorf("got value: %s, want: %s.", actual, expected)
		}
	}
}

func TestFile(t *testing.T) {
	keyValue := map[string]string{
		"path/to/some.yaml": "some",
		"hello.yaml":        "hello",
	}

	for key, value := range keyValue {
		expected := value
		actual := filename(key)
		if expected != actual {
			t.Errorf("got value: %s, want: %s.", actual, expected)
		}
	}
}
