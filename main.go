package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	x, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(x)

	// for _, file := range x {
	// 	fmt.Println(file.Name())
	// }

	//file := 9
	for _, entry := range x {
		if entry.IsDir() {
			fmt.Println(entry.Name(), "is a directory")
		} else {
			fmt.Println(entry.Name(), "is a file")
		}
	}

}
