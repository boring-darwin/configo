package main

import (
	"fmt"

	parse "github.com/boring-darwin/configo"
)

func main() {
	m := parse.ReadConfig("config.ini")
	fmt.Println(m)
}
