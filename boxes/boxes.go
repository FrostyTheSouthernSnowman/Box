package boxes

var Boxes []string

func GetBoxes() *[]string {
	return &Boxes
}

func AddBox(name string) bool {
	Boxes = append(Boxes, name)
	return true
}

type Box struct {
	Begining string
	End      string
}
