package chr

type Controls map[string]string

func (el Controls) Add(name, selector string) {
	el[name] = selector
}

func (el Controls) Find(name string) string {
	return el[name]
}
