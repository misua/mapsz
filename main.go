package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "ff00123",
		"green": "ff2fef",
	}

	// for _, color := range colors {
	// 	fmt.Printf("%+v\n", color)
	// }

	family := map[string]int{
		"dad":  40,
		"mom":  40,
		"sab":  8,
		"chaz": 9,
	}
	printMap(colors)
	printMap(family)
}

// func printMap(c map[string]string) {
// 	for color, hex := range c {
// 		fmt.Printf("%+v - %+v\n", color, hex)
// 	}
// }

func printMap(c interface{}) {
	switch v := c.(type) {
	case map[string]string:
		for key, value := range v {
			fmt.Printf("%+s: %+s\n", key, value)
		}
	case map[string]int:
		for key, value := range v {
			fmt.Printf("%+s: %+v\n", key, value)
		}
	default:
		fmt.Println("unsupported type")
	}
}
