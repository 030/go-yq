package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func scanFile(file string, key string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var val string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key+": ") {
			val = value(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return val
}

func value(keyValue string) string {
	r := regexp.MustCompile("[0-9A-Za-z_]+: ([0-9A-Za-z_:]+).*$")
	match := r.FindStringSubmatch(keyValue)

	if len(match) == 0 {
		log.Fatal("Cannot extract value for key, but was: '", keyValue,
			"' please check whether the regex matches the key")
	}

	return match[1]
}

func dir(path string) string {
	return filepath.Dir(path)
}

func filename(path string) string {
	basename := filepath.Base(path)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	return filepath.Base(filename)
}

func main() {
	// if len(os.Args) <= 2 {
	// 	log.Fatal("Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>")
	// }

	// key := os.Args[1]
	// yamlFile := os.Args[2]

	// value := scanFile(yamlFile, key)
	// if value == "" {
	// 	log.Fatal("File: ", yamlFile, " does not contain key: ", key)
	// }
	// fmt.Println(value)

}
