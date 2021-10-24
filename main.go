package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
)

func FindBox() ([]string, error) {
	var files []string

	rootDir := os.Args[1]

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".box" {
			files = append(files, path)
		}

		return nil
	})

	if err == nil {
		return files, nil
	} else {
		log.Panicf("Could not find file in specified dir: %s", rootDir)
	}

	return nil, nil
}

func ReadBoxFile(filepath string, port string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	box := string(data)
	boxes.MakeBoxFile(box, port)
}

/*func main() {
	fmt.Println("Finding box configuration")
	files, err := FindBox()

	if files == nil && err == nil {
		log.Fatalf("Could not find file in specified dir: %s", os.Args[1])
	} else {
		fmt.Println("Buidling your project")
		ReadBoxFile(files[0], "7000")
	}
}
*/
func main() {
	boxes.ParseYAML("example/box.yml")
}
