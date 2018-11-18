package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type input struct {
	key, file string
}

func (i input) value() string {
	verifyKey(i.key)
	keyWithoutFirstDot := strings.Replace(i.key, ".", "", 1)

	viper.SetConfigType("yaml")

	filename := filename(i.file)
	viper.SetConfigName(filename)

	dir := dir(i.file)
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	value := fmt.Sprintf("%s", viper.Get(keyWithoutFirstDot))

	if value == "%!s(<nil>)" {
		log.Fatal("File: ", i.file, " does not contain key: ", i.key)
	}

	return value
}

func dir(path string) string {
	return filepath.Dir(path)
}

func filename(path string) string {
	basename := filepath.Base(path)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	return filepath.Base(filename)
}

func verifyKey(key string) {
	if !strings.HasPrefix(key, ".") {
		log.Fatal("Key should start with a dot, i.e.: ."+key+", but was: ", key)
	}
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>")
	}

	i := input{key: os.Args[1], file: os.Args[2]}

	fmt.Println(i.value())
}
