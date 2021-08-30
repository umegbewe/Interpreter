package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
