package boxes

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type YAML struct {
	Box struct {
		Port      int    `yaml:"port"`
		Build     string `yaml:"build"`
		Build_bin bool   `yaml:"build-binary"`
	} `yaml:"box"`
}

func ParseYAML(path string) YAML {
	yfile := read_file(path, ".yml")

	var data YAML

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	return data
}

func read_file(path string, ext string) []byte {
	yfile, err := ioutil.ReadFile(path + "/box" + ext)

	if err != nil && ext == ".yml" {
		yaml_file := read_file(path, ".yaml")
		return yaml_file
	} else if err != nil && ext == ".yaml" {
		log.Fatal(err)
	}

	return yfile
}
