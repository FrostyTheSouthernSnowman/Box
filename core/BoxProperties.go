package core

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FrostyTheSouthernSnowman/Box/boxes"
)

func GetStringInBetween(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}

func getProperties(box string) BoxConfig {
	var code []string
	var BoxName string
	var fileExtension string
	var baseBox boxes.Box
	boxes := *boxes.GetBoxes()

	for _, line := range strings.Split(strings.TrimSuffix(box, "\n"), "\n") {
		if line == "" {
			continue
		}

		firstChar := string(line[0])
		comment := "#"

		if firstChar == comment {
			continue
		}

		if strings.HasPrefix(line, "[ ") && strings.HasSuffix(line, " ]") {
			if strings.HasPrefix(line, "[ lang=") {
				if strings.HasSuffix(line, "'python3' ]") {
					fileExtension = ".py"
				} else {
					log.Fatalf("Unsuported language!")
				}
			} else if strings.HasPrefix(line, "[ box_name='") {
				BoxName, _ = GetStringInBetween(line, "[ box_name='", "' ]")
			} else if strings.HasPrefix(line, "[ base_box='") {
				boxName, _ := GetStringInBetween(line, "[ base_box='", "' ]")
				for _, box := range boxes {
					if box.Name == boxName {
						baseBox.Name = boxName
					}
				}
			} else if line == "[ use_default ]" {
				continue
			} else if line == "[ end_box ]" {
				break
			}
		} else {
			code = append(code, line)
		}
	}

	BoxName = BoxName + fileExtension

	BoxProperties := CreateBoxConfig(BoxName, baseBox, code)

	return BoxProperties
}

func MakeBoxFile(dir string, box string, config YAML) {
	var code []string
	properties := getProperties(box)
	filename := config.Box.Build
	if properties.Base.Name == "flask-web-server" {
		box := boxes.FlaskWebServer(strconv.Itoa(config.Box.Port))
		code = append([]string{box.Begining}, properties.Code...)
		code = append(code, box.End)
	}
	if !config.Box.Build_bin {
		NewFile, err := os.Create(dir + "/" + filename)
		if err != nil {
			fmt.Println(err)
			NewFile.Close()
			return
		}

		for _, line := range code {
			fmt.Fprintln(NewFile, line)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		err = NewFile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("finished building your box!")
}
