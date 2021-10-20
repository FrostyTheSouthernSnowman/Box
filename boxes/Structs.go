package boxes

type BoxConfig struct {
	name string
	base string
	code []string
}

func CreateBoxConfig(name string, base string, code []string) BoxConfig {
	return BoxConfig{name: name, base: base, code: code}
}
