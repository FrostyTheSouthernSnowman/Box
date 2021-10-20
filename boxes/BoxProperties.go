package boxes

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getStringInBetween(str string, startS string, endS string) (result string, found bool) {
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
	var baseBox string

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
				BoxName, _ = getStringInBetween(line, "[ box_name='", "' ]")
			} else if strings.HasPrefix(line, "[ base_box='") {
				if line == "[ base_box='flask-web-server' ]" {
					baseBox = "flask-web-server"
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

func MakeBoxFile(box string, port string) {
	var code []string
	properties := getProperties(box)
	filename := properties.name
	if properties.base == "flask-web-server" {
		begining, end := FlaskWebServer(port)
		code = append([]string{begining}, properties.code...)
		code = append(code, end)
	}
	NewFile, err := os.Create(os.Args[1] + "/" + filename)
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
	fmt.Println("finished building your box!")
}
