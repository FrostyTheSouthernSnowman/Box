package core

type BoxConfig struct {
	Name string
	Base string
	Code []string
}

func CreateBoxConfig(name string, base string, code []string) BoxConfig {
	return BoxConfig{Name: name, Base: base, Code: code}
}
