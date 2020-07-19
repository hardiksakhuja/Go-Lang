package main

import "fmt"

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	fmt.Println(string(ret))
}

func cipher(r rune, delta int) rune {
	if r > 'a' && r < 'z' {
		return rotate(r, 'a', delta)
	}
	if r > 'A' && r < 'Z' {
		return rotate(r, 'A', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	temp := int(r) - base
	temp = (temp + delta) % 26
	return rune(temp + base)
}
