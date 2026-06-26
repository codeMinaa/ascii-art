package main

import (
	"fmt"
)

func Render(text string, banner []string) {
	if text == "" {
        return
	}
	for row := 0; row < 8; row++ {
		for _, ch := range text {
			start := (int(ch - 32)) * 9
			fmt.Print(banner[start+row])
		}
		fmt.Println()
	}
}
