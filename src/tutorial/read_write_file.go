package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() (string, error) {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	return string(data), err
}

func writeFile(value []byte) {
	err := ioutil.WriteFile("write_file.txt", value, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func writeFileAtEnd(value string) {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if _, err = file.WriteString(value);err != nil {
		panic(err)
	}
}

func main() {
	output, _ := readFile()
	fmt.Println(output)
	data := "\nThis the writing file at the end of the file"
	value := []byte(data)
	writeFile(value)
	writeFileAtEnd(data)
	res, _ := readFile()
	fmt.Println(res)
}
