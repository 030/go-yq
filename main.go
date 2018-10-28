package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func scanFile(key string) string {
	f, err := os.Open("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var val string

	for scanner.Scan() {
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
	re := regexp.MustCompile("^[a-z]+: (.*)$")
	match := re.FindStringSubmatch(keyValue)
	return match[1]
}

// func readYAML() string {
// 	bytes, err := ioutil.ReadFile("test.yaml")
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	str := string(bytes)

// 	return str
// }

func main() {
	fmt.Println(scanFile("a"))
}
