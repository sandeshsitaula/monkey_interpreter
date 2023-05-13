package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sandeshsitaula/monkeyinter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is monkey language", user.Username)
	fmt.Printf("Type some commands")
	repl.Start(os.Stdin, os.Stdout)
}
