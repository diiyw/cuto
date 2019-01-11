package layertree

import (

	"github.com/diiyw/goc/protocol/dom"

)

// Unique Layer identifier.
type LayerId string	

// Unique snapshot identifier.
type SnapshotId string	

// Rectangle where scrolling happens on the main thread.
type ScrollRect struct {
	
	// Rectangle itself.
	
	Rect	dom.Rect	`json:"rect"`
	
	// Reason for rectangle to force scrolling on the main thread
	// Possible value: RepaintsOnScroll,TouchEventHandler,WheelEventHandler,
	Type	string	`json:"type"`
	
}	

// Sticky position constraints.
type StickyPositionConstraint struct {
	
	// Layout rectangle of the sticky element before being shifted
	
	StickyBoxRect	dom.Rect	`json:"stickyBoxRect"`
	
	// Layout rectangle of the containing block of the sticky element
	
	ContainingBlockRect	dom.Rect	`json:"containingBlockRect"`
	
	// The nearest sticky layer that shifts the sticky box
	
	NearestLayerShiftingStickyBox	LayerId	`json:"nearestLayerShiftingStickyBox"`
	
	// The nearest sticky layer that shifts the containing block
	
	NearestLayerShiftingContainingBlock	LayerId	`json:"nearestLayerShiftingContainingBlock"`
	
}	

// Serialized fragment of layer picture along with its offset within the layer.
type PictureTile struct {
	
	// Offset from owning layer left boundary
	
	X	float64	`json:"x"`
	
	// Offset from owning layer top boundary
	
	Y	float64	`json:"y"`
	
	// Base64-encoded snapshot data.
	
	Picture	string	`json:"picture"`
	
}	

// Information about a compositing layer.
type Layer struct {
	
	// The unique id for this layer.
	
	LayerId	LayerId	`json:"layerId"`
	
	// The id of parent (not present for root).
	
	ParentLayerId	LayerId	`json:"parentLayerId"`
	
	// The backend id for the node associated with this layer.
	
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId"`
	
	// Offset from parent layer, X coordinate.
	
	OffsetX	float64	`json:"offsetX"`
	
	// Offset from parent layer, Y coordinate.
	
	OffsetY	float64	`json:"offsetY"`
	
	// Layer width.
	
	Width	float64	`json:"width"`
	
	// Layer height.
	
	Height	float64	`json:"height"`
	
	// Transformation matrix for layer, default is identity matrix
	
	Transform	[]float64	`json:"transform"`
	
	// Transform anchor point X, absent if no transform specified
	
	AnchorX	float64	`json:"anchorX"`
	
	// Transform anchor point Y, absent if no transform specified
	
	AnchorY	float64	`json:"anchorY"`
	
	// Transform anchor point Z, absent if no transform specified
	
	AnchorZ	float64	`json:"anchorZ"`
	
	// Indicates how many time this layer has painted.
	
	PaintCount	int	`json:"paintCount"`
	
	// Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	
	DrawsContent	bool	`json:"drawsContent"`
	
	// Set if layer is not visible.
	
	Invisible	bool	`json:"invisible"`
	
	// Rectangles scrolling on main thread only.
	
	ScrollRects	[]ScrollRect	`json:"scrollRects"`
	
	// Sticky position constraint information
	
	StickyPositionConstraint	StickyPositionConstraint	`json:"stickyPositionConstraint"`
	
}	

// Array of timings, one per paint step.
type PaintProfile []float64	

