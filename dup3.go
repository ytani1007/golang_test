package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main(){
	counts := make(map[string]int)
	files  := make(map[string]int)
	for i, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		for _, line := range strings.Split(string(data),"\n"){
			counts[line]++
			files[line] += ( i * 2 + 1 )
		}
	}		
	for text, value := range counts {
		if value > 1 {
			fmt.Printf("%d\t%s\t|%d\n", value, text, files[text]);
		}
	}
}
