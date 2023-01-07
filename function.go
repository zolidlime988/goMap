package main

import "fmt"

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Printf("key : %v, value : %v\n", key, val)
	}
}
