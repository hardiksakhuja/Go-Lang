/*
Filename can be passed as a flag. By default filename is passage.txt
passage.txt contains the passage to be processed.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	filename := flag.String("file", "passage.txt", "the input file")
	flag.Parse()
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error Reading File . Error :", err)
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(string(data), " ")
	wordmap := make(map[string]int)
	for _, word := range strings.Fields(processedString) {
		if _, ok := wordmap[word]; ok {
			wordmap[word] = wordmap[word] + 1
		} else {
			wordmap[word] = 1
		}
	}
	total, unique := 0, 0
	for _, v := range wordmap {
		if v == 1 {
			unique++
		}
		total += v
	}
	fmt.Printf("Total Count:%d\n Unique Words: %d", total, unique)
}
