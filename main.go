package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello %s! This is the Monkey proguramming language!\n",
		user.Username)

	fmt.Printf("Feel to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
