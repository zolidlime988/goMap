package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors["white"] = "#fffff"
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	delete(colors, "white")
	fmt.Println(colors)
}
