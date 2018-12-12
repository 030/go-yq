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

func (i input) config() {
	i.verifyKey()

	viper.SetConfigType("yaml")
	viper.SetConfigName(i.filename())
	viper.AddConfigPath(i.dir())
}

func (i input) value() string {
	i.config()
	keyWithoutFirstDot := strings.Replace(i.key, ".", "", 1)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	value := fmt.Sprintf("%v", viper.Get(keyWithoutFirstDot))

	if value == "%!s(<nil>)" {
		log.Fatal("File: ", i.file, " does not contain key: ", i.key)
	}

	return value
}

func (i input) dir() string {
	return filepath.Dir(i.file)
}

func (i input) filename() string {
	basename := filepath.Base(i.file)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	return filepath.Base(filename)
}

func (i input) verifyKey() {
	if !strings.HasPrefix(i.key, ".") {
		log.Fatal("Key should start with a dot, i.e.: ."+i.key+", but was: ", i.key)
	}
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>")
	}

	i := input{key: os.Args[1], file: os.Args[2]}
	fmt.Println(i.value())
}
