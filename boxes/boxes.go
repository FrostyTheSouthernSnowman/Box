package boxes

import "log"

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

func CreateBox(name string) Box {
	var BoxInstance Box
	err := true

	for _, box := range Boxes {
		if box.Name == name {
			BoxInstance = box
			err = false
		}
	}
	if err {
		log.Fatalf("Could not find box with name %s!", name)
	}
	return BoxInstance
}
