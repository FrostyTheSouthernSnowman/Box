package main
 
import (
    "fmt"
    "io/ioutil"
    "strings"
    "log"
    "os"
    "path/filepath"
    "FlaskWebServer"
)

func GetStringInBetween(str string, startS string, endS string) (result string,found bool) {
    s := strings.Index(str, startS)
    if s == -1 {
        return result,false
    }
    newS := str[s+len(startS):]
    e := strings.Index(newS, endS)
    if e == -1 {
        return result,false
    }
    result = newS[:e]
    return result,true
}

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
    filename, file := ParseBox(box, port)
    TranspileBoxFile(filename, file)
}

func ParseBox(BoxConfig string, port string) (string, []string) {
    var code []string
    var fileName string
    var fileExtension string
    var baseBox string

    for _, line := range strings.Split(strings.TrimSuffix(BoxConfig, "\n"), "\n") {
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
                fileName, _ = GetStringInBetween(line, "[ box_name='", "' ]")
            } else if strings.HasPrefix(line, "[ base_box='") {
                if line == "[ base_box='flask-web-server' ]" {
                    begining, _ := FlaskWebServer.FlaskWebServer(port)
                    code = append(code, begining)
                    baseBox = "flask-web-server"
                }
            } else if line == "[ use_default ]" {
                continue
            } else if line == "[ end_box ]" {
                if baseBox == "flask-web-server" {
                    _, end := FlaskWebServer.FlaskWebServer(port)
                    code = append(code, end)
                }
            }
        } else {
            code = append(code, line)
        }
    }
    return fileName + fileExtension, code
}

func TranspileBoxFile(filename string, file []string) {
    NewFile, err := os.Create(os.Args[1] + "/" + filename)
    if err != nil {
        fmt.Println(err)
        NewFile.Close()
        return
    }

    for _, line := range file {
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

/*func CreateFile() {
    fmt.Printf("Writing to a file in Go lang\n")
     
    file, err := os.Create("example/example-box.py")
     
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }

    defer file.Close()

    _, err := file.WriteString("Welcome to GeeksforGeeks."+
            " This program demonstrates reading and writing"+
                         " operations to a file in Go lang.")
 
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
}*/
 
func main() {
    fmt.Println("Finding box configuration")
    files, err := FindBox()

    if files == nil && err == nil{
        log.Fatalf("Could not find file in specified dir: %s", os.Args[1])
    } else {
        fmt.Println("Buidling your project")
        ReadBoxFile(files[0], "7000")
    }
}