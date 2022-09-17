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
	}

	for key, val := range colors {
		fmt.Println(key, val)
	}

	// add a key to a map
	colors["white"] = "#ffffff"

	// delete a key from a map
	delete(colors, "red")

	fmt.Println(colors)
}
