package boxes

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type box struct {
	port         string
	build        string
	build_binary bool
}

type YAML struct {
	versiom string
	boxes   []box
}

func ParseYAML(fileName string) {
	yfile, err := ioutil.ReadFile(fileName)

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[string]YAML)

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	for k, v := range data {

		fmt.Printf("%s: %s\n", k, v)
	}
}
