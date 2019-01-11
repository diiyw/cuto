package frame

// Page frameId
type FrameId string

// Viewport for capturing screenshot.
type Viewport struct {
	
	// X offset in CSS pixels.
	
	X	float64	`json:"x"`
	
	// Y offset in CSS pixels
	
	Y	float64	`json:"y"`

	// Rectangle width in CSS pixels

	Width	float64	`json:"width"`

	// Rectangle height in CSS pixels

	Height	float64	`json:"height"`

	// Page scale factor.

	Scale	float64	`json:"scale"`
}
