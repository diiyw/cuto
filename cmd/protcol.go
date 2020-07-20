package main

import (
	"bytes"
	"strings"
)

type Protocol struct {
	Version struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"version"`
	Domains []Domain
}

type Domain struct {
	Domain       string    `json:"domain"`
	Experimental bool      `json:"experimental,omitempty"`
	Types        []Type    `json:"types,omitempty"`
	Commands     []Command `json:"commands"`
	Events       []Event   `json:"events,omitempty"`
	Description  string    `json:"description,omitempty"`
	Dependencies []string  `json:"dependencies,omitempty"`
}

type Imports []string

func (im Imports) String(domain string) string {
	if len(im) == 0 {
		return ""
	}
	var deps = make(map[string]bool, 0)
	for _, s := range im {
		if s == domain {
			continue
		}
		deps[s] = true
	}
	var buf strings.Builder
	buf.WriteString("\n\nimport (\n")
	for s, _ := range deps {
		buf.WriteString(`	"github.com/diiyw/cuto/protocol/`)
		buf.WriteString(s)
		buf.WriteString("\"\n")
	}
	buf.WriteString(")\n")
	return buf.String()
}

func (d Domain) AllTypes() []byte {
	var pkg bytes.Buffer
	pkg.WriteString("package ")
	domain := strings.ToLower(d.Domain)
	pkg.WriteString(domain)
	var buf strings.Builder
	var imports = make([]string, 0)
	for _, c := range d.Types {
		buf.WriteString("\n")
		deps, str := c.String(domain)
		imports = append(imports, deps...)
		buf.WriteString(str)
	}
	pkg.WriteString(Imports(imports).String(domain))
	pkg.WriteString(buf.String())
	return pkg.Bytes()
}

func (d Domain) AllMethods() []byte {
	var pkg bytes.Buffer
	pkg.WriteString("package ")
	domain := strings.ToLower(d.Domain)
	pkg.WriteString(domain)
	var buf strings.Builder
	var imports = make([]string, 0)
	for _, c := range d.Commands {
		buf.WriteString("\n")
		deps, str := c.String(d.Domain)
		imports = append(imports, deps...)
		buf.WriteString(str)
	}
	pkg.WriteString(Imports(imports).String(domain))
	pkg.WriteString(buf.String())
	return pkg.Bytes()
}

func (d Domain) AllEvents() []byte {
	var pkg bytes.Buffer
	pkg.WriteString("package ")
	domain := strings.ToLower(d.Domain)
	pkg.WriteString(domain)
	var buf strings.Builder
	var imports = make([]string, 0)
	for _, c := range d.Events {
		buf.WriteString("\n")
		deps, str := c.String(d.Domain)
		imports = append(imports, deps...)
		buf.WriteString(str)
	}
	pkg.WriteString(Imports(imports).String(domain))
	pkg.WriteString(buf.String())
	return pkg.Bytes()
}

type Type struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Properties  []Parameter `json:"properties,omitempty"`
	Items       Items       `json:"items,omitempty"`
	Description string      `json:"description"`
}

func (t Type) String(domain string) ([]string, string) {
	var buf strings.Builder
	var imports = make([]string, 0)
	comment := strings.Replace(t.Description, "\n", "\n	// ", -1)
	buf.WriteString(`// ` + comment + "\ntype ")
	buf.WriteString(t.Id)
	dep, typeString := t.genType(domain, t.Id)
	imports = append(imports, dep...)
	buf.WriteString(" " + typeString)
	buf.WriteString("\n")
	return imports, buf.String()
}

func (t Type) genType(domain string, typeID string) ([]string, string) {
	var buf strings.Builder
	var imports = make([]string, 0)
	switch t.Type {
	case "string":
		buf.WriteString("string")
	case "number":
		buf.WriteString("float64")
	case "integer":
		buf.WriteString("int")
	case "any":
		buf.WriteString("interface{}")
	case "binary":
		buf.WriteString("[]byte")
	case "object":
		if len(t.Properties) != 0 {
			buf.WriteString(` struct {`)
			buf.WriteString("\n")
			for _, param := range t.Properties {
				buf.WriteString("\n")
				buf.WriteString(`	// ` + strings.Replace(param.Description, "\n", "\n	// ", -1) + "\n")
				buf.WriteString("	")
				buf.WriteString(strings.ToUpper(param.Name[:1]))
				buf.WriteString(param.Name[1:])
				deps, str := param.genType(domain, typeID)
				buf.WriteString(str)
				buf.WriteString("	`json:\"")
				buf.WriteString(param.Name)
				if param.Optional {
					buf.WriteString(",omitempty")
				}
				buf.WriteString("\"`")
				imports = append(imports, deps...)
				buf.WriteString("\n")
			}
			buf.WriteString("}")
			break
		}
		buf.WriteString("interface{}")
	case "array":
		deps, items := t.Items.String()
		imports = append(imports, deps...)
		buf.WriteString(items)
	}
	return imports, buf.String()
}

type Command struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	Returns     []Parameter `json:"returns,omitempty"`
}

func (c Command) String(domain string) ([]string, string) {
	var buf strings.Builder
	buf.WriteString("\n")
	comment := strings.Replace(c.Description, "\n", "\n// ", -1)
	buf.WriteString(`// ` + comment + "\n")
	buf.WriteString(`const `)
	typeName := strings.ToUpper(c.Name[:1]) + c.Name[1:]
	buf.WriteString(typeName)
	buf.WriteString(" = \"")
	buf.WriteString(domain)
	buf.WriteString(".")
	buf.WriteString(c.Name + "\"\n\n")
	buf.WriteString(`type `)
	buf.WriteString(typeName)
	buf.WriteString(`Params struct {`)
	buf.WriteString("\n")
	// params
	var imports = make([]string, 0)
	for _, param := range c.Parameters {
		dep, paramString := param.String(domain, "")
		imports = append(imports, dep...)
		buf.WriteString(paramString)
		buf.WriteString("	`json:\"")
		buf.WriteString(param.Name)
		if param.Optional {
			buf.WriteString(",omitempty")
		}
		buf.WriteString("\"`\n")
	}
	buf.WriteString("}\n\n")
	buf.WriteString(`type `)
	buf.WriteString(typeName)
	buf.WriteString(`Result`)
	buf.WriteString(" struct {\n")
	// result
	for _, param := range c.Returns {
		dep, paramString := param.String(domain, "")
		imports = append(imports, dep...)
		buf.WriteString(paramString)
		buf.WriteString("	`json:\"")
		buf.WriteString(param.Name)
		buf.WriteString("\"`")
	}
	buf.WriteString("\n")
	buf.WriteString("}")
	return imports, buf.String()
}

type Event struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

func (e Event) String(domain string) ([]string, string) {
	var buf strings.Builder
	buf.WriteString("\n")
	comment := strings.Replace(e.Description, "\n", "\n// ", -1)
	buf.WriteString(`// ` + comment + "\n")
	buf.WriteString(`const `)
	name := strings.ToUpper(e.Name[:1]) + e.Name[1:]
	buf.WriteString(name)
	buf.WriteString(`Event = "`)
	buf.WriteString(domain)
	buf.WriteString(".")
	buf.WriteString(e.Name)
	buf.WriteString("\"\n")
	buf.WriteString("type ")
	buf.WriteString(name)
	buf.WriteString("Params struct {\n")
	// params
	var imports = make([]string, 0)
	for _, param := range e.Parameters {
		dep, paramString := param.String(domain, "")
		imports = append(imports, dep...)
		buf.WriteString(paramString)
	}
	buf.WriteString("}\n\n")
	return imports, buf.String()
}

type Items struct {
	Ref  string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
}

func (items Items) String() ([]string, string) {
	var buf strings.Builder
	var imports = make([]string, 0)
	buf.WriteString("	[]")
	if items.Ref != "" {
		buf.WriteString("*")
		ref := strings.Split(items.Ref, ".")
		if len(ref) == 2 {
			pkg := strings.ToLower(ref[0])
			imports = append(imports, pkg)
			buf.WriteString(pkg)
			buf.WriteString(".")
			buf.WriteString(ref[1])
		} else {
			buf.WriteString(ref[0])
		}
	}
	switch items.Type {
	case "string":
		buf.WriteString("string")
	case "number":
		buf.WriteString("float64")
	case "integer":
		buf.WriteString("int")
	case "any", "object":
		buf.WriteString("interface{}")
	}
	return imports, buf.String()
}

type Parameter struct {
	Name        string   `json:"name"`
	Type        string   `json:"type,omitempty"`
	Ref         string   `json:"$ref,omitempty"`
	Items       Items    `json:"items,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	Description string   `json:"description"`
}

func (param Parameter) genType(domain string, typeID string) ([]string, string) {
	var buf strings.Builder
	var imports = make([]string, 0)
	switch param.Type {
	case "string":
		buf.WriteString("	string")
	case "number":
		buf.WriteString("	float64")
	case "integer":
		buf.WriteString("	int")
	case "any", "object":
		buf.WriteString("	interface{}")
	case "boolean":
		buf.WriteString("	bool")
	case "binary":
		buf.WriteString("	[]byte")
	case "array":
		dep, items := param.Items.String()
		imports = append(imports, dep...)
		buf.WriteString(items)
	default:
		buf.WriteString("	")
		if param.Ref != "" {

			param.Ref = strings.Replace(param.Ref, "DOM.Rect", "cdp.Rect", -1)
			param.Ref = strings.Replace(param.Ref, "DOM.RGBA", "cdp.RGBA", -1)
			param.Ref = strings.Replace(param.Ref, "Page.FrameId", "cdp.FrameId", -1)
			param.Ref = strings.Replace(param.Ref, "Page.Viewport", "cdp.Viewport", -1)
			param.Ref = strings.Replace(param.Ref, "Network.TimeSinceEpoch", "cdp.TimeSinceEpoch", -1)

			param.Ref = strings.Replace(param.Ref, domain+".", "", -1)
			if typeID == param.Ref {
				buf.WriteString("*")
			}
			ref := strings.Split(param.Ref, ".")
			if len(ref) == 2 {
				im := strings.ToLower(ref[0])
				buf.WriteString(im)
				imports = append(imports, im)
				buf.WriteString(".")
				buf.WriteString(ref[1])
			} else {
				buf.WriteString(ref[0])
			}
			break
		}
		buf.WriteString("interface{}")
	}
	return imports, buf.String()
}

func (param Parameter) String(domain, typeID string) ([]string, string) {
	var buf strings.Builder
	buf.WriteString("\n")
	comment := strings.Replace(param.Description, "\n", "\n	// ", -1)
	buf.WriteString(`	// ` + comment + "\n")
	buf.WriteString(`	`)
	buf.WriteString(strings.ToUpper(param.Name[:1]))
	buf.WriteString(param.Name[1:])
	imports, types := param.genType(domain, typeID)
	buf.WriteString(" " + types)
	return imports, buf.String()
}
