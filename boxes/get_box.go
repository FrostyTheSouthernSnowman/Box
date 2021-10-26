package boxes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FindBox(rootDir string) ([]string, error) {
	var files []string

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

func ReadBoxFile(dir string, filepath string, yaml_config YAML) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	box := string(data)
	MakeBoxFile(dir, box, yaml_config)
}
