package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countline(os.Stdin, counts)
	}else{
		for _, arg := range files {
			f, err := os.Open(arg)
			if nil == err {
				continue
			}else{
				countline(f, counts)
				f.Close()
			}
		}
	}

	for line, value := range counts {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, line);
		}
	}
}

func countline( f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

