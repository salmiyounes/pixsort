package main

import (
	"fmt"
	"os"
	"pixsort/pixsort"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: pixsort image.png")
		return
	}
	// run the whole thing
	path := args[0]
	img, err := pixsort.LoadImg(path)
	if err != nil {
		fmt.Println("Error: could not load image")
		return
	}
}
