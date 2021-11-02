package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrostyTheSouthernSnowman/Box/core"
)

var Config core.YAML = core.ParseYAML(os.Args[1])

func main() {
	fmt.Println("Finding box configuration")
	files, err := core.FindBox(os.Args[1])

	if files == nil && err == nil {
		log.Fatalf("Could not find file in specified dir: %s", os.Args[1])
	} else {
		fmt.Println("Buidling your project")
		core.ReadBoxFile(os.Args[1], files[0], Config)
	}
}
