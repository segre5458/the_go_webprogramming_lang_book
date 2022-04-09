package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	counts := make(map[string]int)
	file := make(map[string][]string,0)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, file)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, file)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line,file[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		counts[txt]++
		file[txt] = append(file[txt],f.Name())
	}
}
