package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.ReadFile("./wassup.txt")
	check(err)

	stream, err := tokenize(file)

	if err != nil {
		panic(err)
	}

	stream[0].funcPrint()
	stream[1].varPrint()
}
