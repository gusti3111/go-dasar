package main

import (
	"fmt"
)

func main() {

	m := make(map[string]string)
	m["eat"] = "makan"
	m["drink"] = "minum"
	m["sleep"] = "tidur"
	for e, i := range m {
		fmt.Printf("%s is %s\n", e, i)
	}
	fmt.Println("map:", m)
	// 2
	a := map[int]string{
		1: "satu", 2: "dua", 3: "tiga",
	}
	delete(a, 1)
	fmt.Println("map:", a)

}
