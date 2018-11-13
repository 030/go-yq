package main

import "testing"

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

func TestMultiLineToSingleKey(t *testing.T) {
	keyValue := map[string]string{
		"foo: a": ".foo",
		// "foo: a\nboo: c": ".foo",
		"foo:\n  bar: b":         ".foo.bar",
		"foo:\n  bar:\n  boo: c": ".foo.bar.boo",
	}

	for key, value := range keyValue {
		expected := value
		actual := multiLineToSingleKey(key)
		if expected != actual {
			t.Errorf("Expected %s, but was %s", expected, actual)
		}
	}
}
