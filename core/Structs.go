package core

import "github.com/FrostyTheSouthernSnowman/Box/boxes"

type BoxConfig struct {
	Name string
	Base boxes.Box
	Code []string
}

func CreateBoxConfig(name string, base boxes.Box, code []string) BoxConfig {
	return BoxConfig{Name: name, Base: base, Code: code}
}
