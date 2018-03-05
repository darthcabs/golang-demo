package main

import "fmt"

func main() {
	runes := "א説り자"

	for index, runeValue := range runes {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Printf("%d bytes\n", len(runes))
}
