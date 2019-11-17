package main

import "os/user"

import "fmt"

import "phour/repl"

import "os"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!. This is Phour the programming language!\n", user.Username)
	fmt.Printf("Write this shiny REPL out\n")
	repl.Start(os.Stdin, os.Stdout)
}
