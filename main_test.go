package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

const (
	testYamlFilename = "test2"
	testYaml         = testYamlFilename + ".yaml"
)

var data = `
a: abc
b: def
c: ghi
under_scores: ensureThatKeysMayContainAnUnderscore
services:
  db:
    image: someimage
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: somewordpress
thiscontainsadigit1: helloworld1
alloallo: hallohallo  # yamllint disable-line rule:line-length
foo:
  bar: boo
firefox_checksum: sha512:49d776
hello:
  world: hallo wereld
world: [hola, hallo]
firefox_version4: 64
firefox_version3: 64.1
firefox_version2: "64.0"
firefox_version: 64.0.0
  `

// See https://stackoverflow.com/a/34102842/2777965
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	createTestYaml()
}

func shutdown() {
	removeTestYaml()
}

func createTestYaml() {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(testYaml, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func removeTestYaml() {
	os.Remove(testYaml)
}

func TestYamlValue(t *testing.T) {
	createTestYaml()

	keyValue := map[string]string{
		".a":                   "abc",
		".b":                   "def",
		".c":                   "ghi",
		".under_scores":        "ensureThatKeysMayContainAnUnderscore",
		".thiscontainsadigit1": "helloworld1",
		".alloallo":            "hallohallo",
		".firefox_checksum":    "sha512:49d776",
		".foo.bar":             "boo",
		".services.db.image":   "someimage",
		".services.db.environment.MYSQL_ROOT_PASSWORD": "somewordpress",
		".world":            "[hola hallo]",
		".hello.world":      "hallo wereld",
		".firefox_version4": "64",
		".firefox_version3": "64.1",
		".firefox_version2": "64.0",
		".firefox_version":  "64.0.0",
	}

	for key, value := range keyValue {
		i := input{key: key, file: testYaml}

		expected := value
		actual, _ := i.value()
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

func TestVerifyKey(t *testing.T) {
	i := input{"abc", testYamlFilename}

	err := i.verifyKey()
	want := "Key should start with a dot, i.e.: .abc, but was: abc"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}
}

func TestValue(t *testing.T) {
	i := input{".abc", testYaml}

	_, err := i.value()
	want := "File: test2.yaml does not contain key: .abc"
	if err.Error() != want {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}
}

func TestReadInConfig(t *testing.T) {
	i := input{".abc", "fileDoesNotExist"}

	_, err := i.value()
	want := "fatal error config file: Config File \"fileDoesNotExist\" Not Found in"
	matched, _ := regexp.MatchString(want, err.Error())
	if !matched {
		t.Errorf("Error expected. Got '%v'. Want '%v'", err, want)
	}
}
