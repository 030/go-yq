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
	viper.SetConfigType("yaml")
	viper.SetConfigName(i.filename())
	viper.AddConfigPath(i.dir())
}

func (i input) value() (string, error) {
	i.config()
	keyWithoutFirstDot := strings.Replace(i.key, ".", "", 1)

	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("fatal error config file: %v", err)
	}
	value := fmt.Sprintf("%v", viper.Get(keyWithoutFirstDot))

	if value == "<nil>" {
		return "", fmt.Errorf("File: %v does not contain key: %v", i.file, i.key)
	}

	return value, nil
}

func (i input) dir() string {
	return filepath.Dir(i.file)
}

func (i input) filename() string {
	basename := filepath.Base(i.file)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))

	return filepath.Base(filename)
}

func (i input) verifyKey() error {
	if !strings.HasPrefix(i.key, ".") {
		return fmt.Errorf("Key should start with a dot, i.e.: .%s, but was: %s", i.key, i.key)
	}
	return nil
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Usage: go-yq <key e.g. .foo.bar> <filename e.g. input.yaml>")
	}

	i := input{key: os.Args[1], file: os.Args[2]}

	err := i.verifyKey()
	if err != nil {
		log.Fatal(err)
	}

	v, err := i.value()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
