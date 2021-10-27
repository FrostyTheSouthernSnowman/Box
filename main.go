package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Finding box configuration")
	files, err := core.FindBox(os.Args[1])
	config := core.ParseYAML(os.Args[1])

	if files == nil && err == nil {
		log.Fatalf("Could not find file in specified dir: %s", os.Args[1])
	} else {
		fmt.Println("Buidling your project")
		core.ReadBoxFile(os.Args[1], files[0], config)
	}
}
