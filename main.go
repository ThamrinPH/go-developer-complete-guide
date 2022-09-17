package main

import "fmt"

func main() {
	// cara inisialisasi variable bertipe map ada beberapa cara
	// #1
	// var colors map[string]string -> menghasilnya map kosong
	// #2
	// colors := make(map[string]string) -> menghasilnya map kosong

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff4582",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}
