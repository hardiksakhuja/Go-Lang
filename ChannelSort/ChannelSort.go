package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInput() (string, error) {
	stdinReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a sequence of integers separated by space")
	inputString, err := stdinReader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input. Error:", err)
		return "", err
	} else {
		if strings.HasSuffix(inputString, "\n") {
			inputString = strings.TrimSuffix(inputString, "\n")
		}

		if strings.HasSuffix(inputString, "\r") {
			inputString = strings.TrimSuffix(inputString, "\r")
		}

		inputString = strings.TrimSpace(inputString)

		return inputString, nil
	}
}

func toIntegers(str string) ([]int, error) {
	parts := strings.Split(str, " ")
	// fmt.Println(parts)
	var slice []int
	// fmt.Println(slice)
	for _, v := range parts {
		if n, err := strconv.Atoi(v); err != nil {
			fmt.Print("error", err)
			return nil, err
		} else {
			// fmt.Printf("\n %T,%v \n", n, n)
			slice = append(slice, n)
		}
	}
	// fmt.Println(slice)
	return slice, nil
}

func sort_array(arr []int, channel chan []int) {

	orig_arr := make([]int, len(arr))
	copy(orig_arr, arr)
	sort.Ints(arr)
	fmt.Println("Sorted", orig_arr, "into", arr)

	channel <- arr
}

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	var final []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			final = append(final, arr1[i])
			i++
		} else {
			final = append(final, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		for i < len(arr1) {
			final = append(final, arr1[i])
			i++
		}
	}
	if j < len(arr2) {
		for j < len(arr2) {
			final = append(final, arr2[j])
			j++
		}
	}
	return final
}

func main() {
	// slice2 := make([]int, 0)
	// slice2 = append(slice2, 5, 3, 4, 2, 1)
	// fmt.Println(slice2)
	if inputString, err := ReadInput(); err != nil {
		fmt.Println("Error:", err)
	} else {
		arr, _ := toIntegers(inputString)
		// fmt.Println(arr)
		size := len(arr) / 4
		// fmt.Printf("Len %v Size %v", len(arr), size)
		// fmt.Printf("%q\n%q", inputString, arr)
		channel := make(chan []int, 4)
		for c := 0; c < 4; c++ {
			if c != 3 {
				go sort_array(arr[c*size:(c+1)*size], channel)
			} else {
				go sort_array(arr[c*size:], channel)
			}
		}
		sort1 := <-channel
		sort2 := <-channel
		sort3 := <-channel
		sort4 := <-channel
		final := merge(merge(sort1, sort2), merge(sort3, sort4))
		fmt.Println("Final Sorted list:\n", final)
	}
}
