package boxes

var Boxes []Box

func GetBoxes() *[]Box {
	return &Boxes
}

func AddBox(box Box) bool {
	Boxes = append(Boxes, box)
	return true
}

type Box struct {
	Name       string
	Lang       string
	Begining   string
	End        string
	Dockerfile string
}
