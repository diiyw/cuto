package cdp

// Rectangle.
type Rect  struct {

	// X coordinate
	X	float64	`json:"x"`

	// Y coordinate
	Y	float64	`json:"y"`

	// Rectangle width
	Width	float64	`json:"width"`

	// Rectangle height
	Height	float64	`json:"height"`
}

// A structure holding an RGBA color.
type RGBA  struct {

	// The red component, in the [0-255] range.
	R	int	`json:"r"`

	// The green component, in the [0-255] range.
	G	int	`json:"g"`

	// The blue component, in the [0-255] range.
	B	int	`json:"b"`

	// The alpha component, in the [0-1] range (default: 1).
	A	float64	`json:"a"`
}

// Viewport for capturing screenshot.
type Viewport  struct {

	// X offset in device independent pixels (dip).
	X	float64	`json:"x"`

	// Y offset in device independent pixels (dip).
	Y	float64	`json:"y"`

	// Rectangle width in device independent pixels (dip).
	Width	float64	`json:"width"`

	// Rectangle height in device independent pixels (dip).
	Height	float64	`json:"height"`

	// Page scale factor.
	Scale	float64	`json:"scale"`
}


type FrameId string

type TimeSinceEpoch float64
	