package generator

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	frame = strings.Replace(`package frame

// Page frameId
type FrameId string

// Viewport for capturing screenshot.
type Viewport struct {
	
	// X offset in CSS pixels.
	
	X	float64	\x60json:"x"\x60
	
	// Y offset in CSS pixels
	
	Y	float64	\x60json:"y"\x60

	// Rectangle width in CSS pixels

	Width	float64	\x60json:"width"\x60

	// Rectangle height in CSS pixels

	Height	float64	\x60json:"height"\x60

	// Page scale factor.

	Scale	float64	\x60json:"scale"\x60
}
`, `\x60`, "`", -1)

	imports = `package {{$.DomainName|lower}}
{{if .Requires}}
import (
{{range .Requires}}
	"github.com/diiyw/goc/protocol{{.|lower}}"
{{end}}
){{end}}`
	methods = `
const (
	{{range .Commands}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{.Name|uc}} = "{{$.DomainName}}.{{.Name}}"
	{{end}}
)
{{range $v := .Commands}}
// {{.Name|uc}} parameters
type {{.Name|uc}}Params struct {
	{{range .Parameters}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{.Name|uc}}	{{if .Type}}{{getType .Type $v.Name .Items}}{{else}}{{getType .Ref $v.Name .Items}}{{end}}	{{json .Name .Description}}
	{{end}}
}

// {{.Name|uc}} returns
type {{.Name|uc}}Returns struct {
	{{range .Returns}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{.Name|uc}}	{{if .Type}}{{getType .Type $v.Name .Items}}{{else}}{{getType .Ref $v.Name .Items}}{{end}}	{{json .Name .Description}}
	{{end}}
}
{{end}}
`
	events = `
const (
	{{range .Events}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{.Name|uc}}Event = "{{$.DomainName}}.{{.Name}}"
	{{end}}
)
{{range $v := .Events}}
{{if .Description}}// {{.Description|unescaped}}{{end}}
type {{.Name|uc}}Params struct {
	{{range .Parameters}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{.Name|uc}}	{{if .Type}}{{getType .Type $v.Name .Items}}{{else}}{{getType .Ref $v.Name .Items}}{{end}}	{{json .Name .Description}}
	{{end}}
}
{{end}}
`
	types = `
{{range $v := .Types}}
// {{.Description|unescaped}}
type {{.Id}} {{if eq .Type "object"}}struct {
	{{range .Properties}}
	{{if .Description}}// {{.Description|unescaped}}{{end}}
	{{if .Enum}}// Possible value: {{range .Enum}}{{.}},{{end}}{{end}}
	{{.Name|uc}}	{{if .Type}}{{getType .Type $v.Id .Items}}{{else}}{{getType .Ref $v.Id .Items}}{{end}}	{{json .Name .Description}}
	{{end}}
}{{else}}{{getType .Type $v.Id .Items}}{{end}}	
{{end}}
`
)

type Items struct {
	Type string `json:"type,omitempty"`
	Ref  string `json:"ref,omitempty"`
}

type Parameter struct {
	Name        string   `json:"name"`
	Type        string   `json:"type,omitempty"`
	Ref         string   `json:"ref,omitempty"`
	Items       Items    `json:"items,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	Description string   `json:"description"`
}

type Protocol struct {
	Version struct {
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"version"`
	Domains []struct {
		Domain       string `json:"domain"`
		Experimental bool   `json:"experimental,omitempty"`
		Types        []struct {
			Id          string      `json:"id"`
			Type        string      `json:"type"`
			Properties  []Parameter `json:"properties,omitempty"`
			Items       Items       `json:"items,omitempty"`
			Description string      `json:"description"`
		} `json:"types,omitempty"`
		Commands []struct {
			Name        string      `json:"name"`
			Description string      `json:"description"`
			Parameters  []Parameter `json:"parameters,omitempty"`
			Returns     []Parameter `json:"returns,omitempty"`
		} `json:"commands"`
		Events []struct {
			Name        string      `json:"name"`
			Description string      `json:"description"`
			Parameters  []Parameter `json:"parameters,omitempty"`
		} `json:"events,omitempty"`
		Description  string   `json:"description,omitempty"`
		Dependencies []string `json:"dependencies,omitempty"`
	} `json:"domains"`
}

func Generate() error {
	path, _ := os.Getwd()
	path += "/../protocol/"
	base := "https://raw.githubusercontent.com/ChromeDevTools/devtools-protocol/master/json/"
	jsProtocol := base + "js_protocol.json"
	browserProtocol := base + "browser_protocol.json"
	resp, err := http.Get(jsProtocol)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	jsCtn, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp, err = http.Get(browserProtocol)
	if err != nil {
		return err
	}
	browserCtn, err := ioutil.ReadAll(resp.Body)
	jsCtn = bytes.Replace(jsCtn, []byte("$ref"), []byte("ref"), -1)
	browserCtn = bytes.Replace(browserCtn, []byte("$ref"), []byte("ref"), -1)
	var js, browser Protocol
	if err := json.Unmarshal(jsCtn, &js); err != nil {
		return err
	}
	if err := json.Unmarshal(browserCtn, &browser); err != nil {
		return err
	}
	var domains = append(js.Domains, browser.Domains...)
	// 创建Frame模块
	if err := os.MkdirAll(path+"frame", 0777); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path+"frame/types.go", []byte(frame), 0777); err != nil {
		return err
	}
	for _, domain := range domains {
		var requires = make(map[string]string)
		tpl := template.New("template")
		tpl = tpl.Funcs(template.FuncMap{"uc": uc})
		tpl = tpl.Funcs(template.FuncMap{"lower": func(s string) string {
			return strings.ToLower(s)
		}})
		tpl = tpl.Funcs(template.FuncMap{"json": func(s, description string) interface{} {
			if strings.Contains(description,"omitted") {
				s += ",omitempty"
			}
			return template.HTML("`json:\"" + s + "\"`")
		}})
		tpl = tpl.Funcs(template.FuncMap{"unescaped": func(s string) interface{} {
			s = strings.Replace(s, "\n", "\n\t// ", -1)
			return template.HTML(s)
		}})
		tpl = tpl.Funcs(template.FuncMap{"getType": func(s, name string, items Items) string {
			var t = getType(s, uc(name), domain.Domain, items)
			var typeName = t[0]
			if t[1] != "" {
				var dep = strings.TrimLeft(t[1], "[]")
				dep = strings.TrimLeft(dep, "*")
				if t[1] == "[]" {
					typeName = t[1] + t[0]
				} else {
					typeName = t[1] + "." + t[0]
				}
				switch typeName {
				case "page.FrameId":
					typeName = "frame.FrameId"
					requires["/frame"] = "/frame"
				case "page.Viewport":
					typeName = "frame.Viewport"
					requires["/frame"] = "/frame"
				default:
					if dep != "" {
						requires["/"+dep] = "/" + dep
					}
				}
			}
			return typeName
		}})
		domainPath := path + strings.ToLower(domain.Domain)
		err = os.MkdirAll(domainPath, 0777)
		if err != nil {
			return err
		}
		var create = func(s string) error {
			var file, temp = "", ""
			// common data
			data := map[string]interface{}{
				"DomainName":   domain.Domain,
				"Dependencies": domain.Dependencies,
			}
			switch s {
			case "commands":
				temp = methods
				data["Commands"] = domain.Commands
				file = domainPath + "/methods.go"
			case "types":
				temp = types
				data["Types"] = domain.Types
				file = domainPath + "/types.go"
			case "events":
				temp = events
				data["Events"] = domain.Events
				file = domainPath + "/events.go"
			}
			var clone, err = tpl.Clone()
			if err != nil {
				return err
			}
			clone, err = clone.Parse(temp)
			if err != nil {
				return err
			}
			var buf = bytes.NewBuffer(nil)
			err = clone.Execute(buf, data)
			if err != nil {
				return err
			}
			dep, err := tpl.Clone()
			if err != nil {
				return err
			}
			dep, err = dep.Parse(imports)
			if err != nil {
				return err
			}
			fi, err := os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				return err
			}
			data["Requires"] = requires
			if err := dep.Execute(fi, data); err != nil {
				return err
			}
			if _, err = fi.Write(buf.Bytes()); err != nil {
				return err
			}
			_ = fi.Close()
			requires = make(map[string]string)
			return nil
		}
		if err := create("commands"); err != nil {
			return err
		}
		if err := create("types"); err != nil {
			return err
		}
		if err := create("events"); err != nil {
			return err
		}
	}
	return nil
}

func uc(s string) string {
	b := []byte(s)
	f := b[:1]
	f = bytes.ToUpper(f)
	b[0] = f[0]
	return string(b)
}

func getType(s, name, domain string, items Items) (t [2]string) {
	switch s {
	case "string":
		t[0] = s
	case "boolean":
		t[0] = "bool"
	case "integer":
		t[0] = "int"
	case "number":
		t[0] = "float64"
	case "any", "object":
		t[0] = "interface{}"
	case "array":
		if items.Type != "" {
			t := getType(items.Type, name, domain, items)
			t[1] = "[]" + t[1]
			return t
		}
		t := getType(items.Ref, name, domain, items)
		t[1] = "[]" + t[1]
		return t
	default:
		if strings.Contains(s, ".") {
			ss := strings.Split(s, ".")
			if ss[0] == domain {
				t[1] = ""
			} else {
				t[0] = ss[1]
				t[1] = strings.ToLower(ss[0])
			}
		} else {
			if s == name {
				t[0] = "*" + s
			} else {
				t[0] = s
			}

		}
	}
	return t
}
