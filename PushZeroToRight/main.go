package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter space separated integers: ")
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	slice := convert(line)
	fmt.Println("Answer: ", solve(slice))
}

func solve(arr []int) []int {
	count := 0
	for i, v := range arr {
		if v != 0 {
			arr[i], arr[count] = arr[count], arr[i]
			count++
		}
	}
	return arr
}

//to convert array elements to int
func convert(str string) []int {
	arr := make([]int, 0)
	for _, val := range strings.Fields(str) {
		integer, _ := strconv.Atoi(val)
		arr = append(arr, integer)
	}
	return arr
}
