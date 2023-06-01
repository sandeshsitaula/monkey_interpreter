package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/sandeshsitaula/monkeyinter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	var w io.Reader
	/*
		filename := flag.String("f", "main.mn", "go run main.go -f <filename>")
		flag.Parse()
	*/
	if len(os.Args) <= 1 {
		fmt.Printf("Hello %s! This is monkey language", user.Username)
		fmt.Printf("Type some commands")
		w = os.Stdin
	} else {
		filename := os.Args[1]
		if !strings.HasSuffix(filename, ".mn") {
			fmt.Println("Wrong file format. Only  .mn extension allowed")
			os.Exit(1)
		}
		w, err = os.Open(filename)
		if err != nil {
			fmt.Println("No such fille found")
			os.Exit(1)

		}

	}

	repl.Start(w, os.Stdout)
}
