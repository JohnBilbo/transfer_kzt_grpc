package main

import (
	"os"
	"user/internal/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	c, err := config.FindProjectRoot(cwd)
	if err != nil {
		panic(err)
	}
	print(c)
}
