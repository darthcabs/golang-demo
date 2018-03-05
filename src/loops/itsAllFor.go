package main

import "fmt"

func main() {

	// While
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()

	// For clássico
	for j := 4; j <= 6; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	// Do-while
	k := 7
	for {
		fmt.Printf("%d\n", k)

		if k == 9 {
			break
		}
		k++
	}
	fmt.Println()

	for l := 10; l <= 15; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}
}
