package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	input := os.Args[1]
	bannerName := os.Args[2]

	banner, err := LoadBanner("banner/" + bannerName + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	parts := SplitInput(input)

	for i, part := range parts {
		Render(part, banner)

		if i < len(parts)-1 {
			fmt.Println()
		}
	}
}