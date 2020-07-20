package input
// 
type TouchPoint  struct {

	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X	float64	`json:"x"`

	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to
	// the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y	float64	`json:"y"`

	// X radius of the touch area (default: 1.0).
	RadiusX	float64	`json:"radiusX,omitempty"`

	// Y radius of the touch area (default: 1.0).
	RadiusY	float64	`json:"radiusY,omitempty"`

	// Rotation angle (default: 0.0).
	RotationAngle	float64	`json:"rotationAngle,omitempty"`

	// Force (default: 1.0).
	Force	float64	`json:"force,omitempty"`

	// Identifier used to track touch sources between events, must be unique within an event.
	Id	float64	`json:"id,omitempty"`
}

// 
type GestureSourceType string

// UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64
