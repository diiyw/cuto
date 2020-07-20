package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("protocol.json")
	var proto Protocol
	if err := json.Unmarshal(b, &proto); err != nil {
		panic(err)
	}
	var protocol = "../protocol/"
	_ = os.RemoveAll(protocol)
	_ = os.MkdirAll(protocol+"cdp", 0755)
	_ = ioutil.WriteFile(protocol+"cdp/type.go", []byte(`package cdp

// Rectangle.
type Rect  struct {

	// X coordinate
	X	float64	`+"`json:\"x\"`"+`

	// Y coordinate
	Y	float64`+"	`json:\"y\"`"+`

	// Rectangle width
	Width	float64`+"	`json:\"width\"`"+`

	// Rectangle height
	Height	float64	`+"`json:\"height\"`"+`
}

// A structure holding an RGBA color.
type RGBA  struct {

	// The red component, in the [0-255] range.
	R	int	`+"`json:\"r\"`"+`

	// The green component, in the [0-255] range.
	G	int	`+"`json:\"g\"`"+`

	// The blue component, in the [0-255] range.
	B	int	`+"`json:\"b\"`"+`

	// The alpha component, in the [0-1] range (default: 1).
	A	float64	`+"`json:\"a\"`"+`
}

// Viewport for capturing screenshot.
type Viewport  struct {

	// X offset in device independent pixels (dip).
	X	float64	`+"`json:\"x\"`"+`

	// Y offset in device independent pixels (dip).
	Y	float64	`+"`json:\"y\"`"+`

	// Rectangle width in device independent pixels (dip).
	Width	float64	`+"`json:\"width\"`"+`

	// Rectangle height in device independent pixels (dip).
	Height	float64	`+"`json:\"height\"`"+`

	// Page scale factor.
	Scale	float64	`+"`json:\"scale\"`"+`
}


type FrameId string

type TimeSinceEpoch float64
	`), 0755)
	for _, domain := range proto.Domains {
		dirname := protocol + strings.ToLower(domain.Domain)
		if err := os.MkdirAll(dirname, 0755); err != nil {
			panic(err)
		}
		_ = ioutil.WriteFile(dirname+"/type.go", domain.AllTypes(), 0755)
		_ = ioutil.WriteFile(dirname+"/method.go", domain.AllMethods(), 0755)
		_ = ioutil.WriteFile(dirname+"/event.go", domain.AllEvents(), 0755)
	}
}
