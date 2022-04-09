package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := ""
	for i, val := range os.Args[1:] {
		s = strconv.Itoa(i) + " : " + val
		fmt.Println(s)
	}
}
