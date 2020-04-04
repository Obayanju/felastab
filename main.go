package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/obayanju/felastab/search"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is a JavaScript interpreter!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	search.Start(os.Stdin, os.Stdout)
}
