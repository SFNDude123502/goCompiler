package main

import (
	"log"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.ReadFile("./test.gasm")
	check(err)

	stream, err := tokenizeFile(file)

	if err != nil {
		panic(err)
	}

	parsed, imports := parse(stream)

	for _, imp := range imports {
		parsed = "import  \"" + imp + "\" \n" + parsed
	}
	parsed = "package main\n" + parsed

	os.WriteFile("./build/main.go", []byte(parsed), 0666)
	cmd := exec.Command("go", "build", "build/main.go")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
