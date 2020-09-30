package main

import (
	"io/ioutil"
)

func WriteToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func ReadFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}
