package main

import "fmt"

func main() {
	/*
		kapan menggunakan struct atau map vice versa

		map semua nilainya dan keynya memiliki tipe yang sama sedangkan struct bisa berbeda
	*/
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
