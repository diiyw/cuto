package layertree

import (
	"github.com/diiyw/cuto/protocol/cdp"
	"github.com/diiyw/cuto/protocol/dom"
)

// Unique Layer identifier.
type LayerId string

// Unique snapshot identifier.
type SnapshotId string

// Rectangle where scrolling happens on the main thread.
type ScrollRect  struct {

	// Rectangle itself.
	Rect	cdp.Rect	`json:"rect"`

	// Reason for rectangle to force scrolling on the main thread
	Type	string	`json:"type"`
}

// Sticky position constraints.
type StickyPositionConstraint  struct {

	// Layout rectangle of the sticky element before being shifted
	StickyBoxRect	cdp.Rect	`json:"stickyBoxRect"`

	// Layout rectangle of the containing block of the sticky element
	ContainingBlockRect	cdp.Rect	`json:"containingBlockRect"`

	// The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingStickyBox	LayerId	`json:"nearestLayerShiftingStickyBox,omitempty"`

	// The nearest sticky layer that shifts the containing block
	NearestLayerShiftingContainingBlock	LayerId	`json:"nearestLayerShiftingContainingBlock,omitempty"`
}

// Serialized fragment of layer picture along with its offset within the layer.
type PictureTile  struct {

	// Offset from owning layer left boundary
	X	float64	`json:"x"`

	// Offset from owning layer top boundary
	Y	float64	`json:"y"`

	// Base64-encoded snapshot data.
	Picture	[]byte	`json:"picture"`
}

// Information about a compositing layer.
type Layer  struct {

	// The unique id for this layer.
	LayerId	LayerId	`json:"layerId"`

	// The id of parent (not present for root).
	ParentLayerId	LayerId	`json:"parentLayerId,omitempty"`

	// The backend id for the node associated with this layer.
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId,omitempty"`

	// Offset from parent layer, X coordinate.
	OffsetX	float64	`json:"offsetX"`

	// Offset from parent layer, Y coordinate.
	OffsetY	float64	`json:"offsetY"`

	// Layer width.
	Width	float64	`json:"width"`

	// Layer height.
	Height	float64	`json:"height"`

	// Transformation matrix for layer, default is identity matrix
	Transform	[]float64	`json:"transform,omitempty"`

	// Transform anchor point X, absent if no transform specified
	AnchorX	float64	`json:"anchorX,omitempty"`

	// Transform anchor point Y, absent if no transform specified
	AnchorY	float64	`json:"anchorY,omitempty"`

	// Transform anchor point Z, absent if no transform specified
	AnchorZ	float64	`json:"anchorZ,omitempty"`

	// Indicates how many time this layer has painted.
	PaintCount	int	`json:"paintCount"`

	// Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	DrawsContent	bool	`json:"drawsContent"`

	// Set if layer is not visible.
	Invisible	bool	`json:"invisible,omitempty"`

	// Rectangles scrolling on main thread only.
	ScrollRects	[]*ScrollRect	`json:"scrollRects,omitempty"`

	// Sticky position constraint information
	StickyPositionConstraint	StickyPositionConstraint	`json:"stickyPositionConstraint,omitempty"`
}

// Array of timings, one per paint step.
type PaintProfile 	[]float64
