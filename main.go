package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
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
		logrus.Debug("Check whether file: ", file, " and line: ", scanner.Text(), " contains key: ", key)
		if strings.Contains(scanner.Text(), key+":") {
			val = value(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return val
}

func value(keyValue string) string {
	re := regexp.MustCompile("^[a-z-_]+: (.*)$")
	match := re.FindStringSubmatch(keyValue)

	if len(match) == 0 {
		log.Fatal("Cannot extract value for key, but was: '", keyValue,
			"' please check whether the regex matches the key")
	}

	logrus.Debug("MATCH: ", match)

	return match[1]
}

func main() {
	key := flag.String("key", "key", "Specify the key")
	yamlFile := flag.String("yamlFile", "file.yaml", "Path to a yaml file")
	debug := flag.Bool("debug", false, "Whether debugging should be enabled")

	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debug("key: ", *key)
	logrus.Debug("yamlFile: ", *yamlFile)

	value := scanFile(*yamlFile, *key)
	if value == "" {
		log.Fatal("File: ", *yamlFile, " does not contain key: ", *key)
	}
	fmt.Println(value)
}
