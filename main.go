package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
)

func main() {
	fmt.Println("Finding box configuration")
	files, err := boxes.FindBox(os.Args[1])
	config := boxes.ParseYAML(os.Args[1])

	if files == nil && err == nil {
		log.Fatalf("Could not find file in specified dir: %s", os.Args[1])
	} else {
		fmt.Println("Buidling your project")
		boxes.ReadBoxFile(os.Args[1], files[0], config)
	}
}
