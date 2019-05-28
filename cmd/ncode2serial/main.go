package main

import (
	"fmt"
	"os"

	narou "github.com/ress997/go-narou"
)

func main() {
	fmt.Println(narou.Ncode2Serial(os.Args[1]))
}
