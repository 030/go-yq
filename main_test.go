package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
  `

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
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

	ioutil.WriteFile(testYaml, d, 0644)
}

func removeTestYaml() {
	os.Remove(testYaml)
}

func TestYamlValue(t *testing.T) {
	createTestYaml()

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
		i := input{key: key, file: testYaml}

		expected := value
		actual := i.value()
		if expected != actual {
			t.Errorf("Value was incorrect 'Check whether the key '%s' resides in the test yaml file', got value: %s, want: %s.", key, actual, expected)
		}
	}

	removeTestYaml()
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
