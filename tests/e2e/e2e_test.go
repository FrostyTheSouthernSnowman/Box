package e2e__test

import (
	"io/ioutil"
	"testing"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
)

func Test_e2e_using_example(t *testing.T) {
	files, err := boxes.FindBox("test_configs")
	config := boxes.ParseYAML("test_configs")

	if files == nil && err == nil {
		t.Fatalf("Could not find file in specified dir: ./test_config")
	} else {
		boxes.ReadBoxFile("test_configs", files[0], config)
	}

	data, err := ioutil.ReadFile("test_configs/flask-box.py")
	if err != nil {
		t.Fatalf("The file did not genarage in ./test_config")
	}

	correct_box_data, err2 := ioutil.ReadFile("test_configs/correct_build.txt")
	if err2 != nil {
		t.Fatalf("Could not find a box to compare against in test_config")
	}

	box := string(data)
	correct_box := string(correct_box_data)

	if box != correct_box {
		t.Fatalf("The box was not built correctly")
	}
}
