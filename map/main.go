package main

func main() {
	// Create a MAP

	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Delete a map

	// delete(colors, "white")

	printMap(colors)
}

// Iterate a map
func printMap(c map[string]string) {
	for color, hex := range c {
		println("Hex code for", color, "is", hex)
	}
}
