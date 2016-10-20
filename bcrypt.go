package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	cost := flag.Int("cost", bcrypt.DefaultCost, "The bcrypt cost")
	flag.Parse()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read stdin: %v", err)
		os.Exit(2)
	}

	out, err := bcrypt.GenerateFromPassword(data, *cost)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bcrypt password: %v", err)
		os.Exit(2)
	}
	os.Stdout.Write(out)
}
