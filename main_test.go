package main

import (
	"path/filepath"
	"testing"
)

func TestYamlValue(t *testing.T) {
	keyValue := map[string]string{
		".a":                                           "abc",
		".b":                                           "def",
		".c":                                           "ghi",
		".under_scores":                                "ensureThatKeysMayContainAnUnderscore",
		".thiscontainsadigit1":                         "helloworld1",
		".alloallo":                                    "hallohallo",
		".firefox_checksum":                            "sha512:49d776",
		".foo.bar":                                     "boo",
		".services.db.image":                           "someimage",
		".services.db.environment.MYSQL_ROOT_PASSWORD": "somewordpress",
		".world":       "[hola hallo]",
		".hello.world": "hallo wereld",
	}

	for key, value := range keyValue {
		i := input{key: key, file: "test"}

		expected := value
		actual := i.value()
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
	for file, value := range keyValue {
		i := input{key: "", file: file}

		expected := value
		actual := i.dir()
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

	for file, value := range keyValue {
		i := input{key: "", file: file}

		expected := value
		actual := i.filename()
		if expected != actual {
			t.Errorf("got value: %s, want: %s.", actual, expected)
		}
	}
}
