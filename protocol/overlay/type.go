package overlay

import (
	"github.com/diiyw/cuto/protocol/cdp"
)

// Configuration data for the highlighting of page elements.
type HighlightConfig  struct {

	// Whether the node info tooltip should be shown (default: false).
	ShowInfo	bool	`json:"showInfo,omitempty"`

	// Whether the node styles in the tooltip (default: false).
	ShowStyles	bool	`json:"showStyles,omitempty"`

	// Whether the rulers should be shown (default: false).
	ShowRulers	bool	`json:"showRulers,omitempty"`

	// Whether the extension lines from node to the rulers should be shown (default: false).
	ShowExtensionLines	bool	`json:"showExtensionLines,omitempty"`

	// The content box highlight fill color (default: transparent).
	ContentColor	cdp.RGBA	`json:"contentColor,omitempty"`

	// The padding highlight fill color (default: transparent).
	PaddingColor	cdp.RGBA	`json:"paddingColor,omitempty"`

	// The border highlight fill color (default: transparent).
	BorderColor	cdp.RGBA	`json:"borderColor,omitempty"`

	// The margin highlight fill color (default: transparent).
	MarginColor	cdp.RGBA	`json:"marginColor,omitempty"`

	// The event target element highlight fill color (default: transparent).
	EventTargetColor	cdp.RGBA	`json:"eventTargetColor,omitempty"`

	// The shape outside fill color (default: transparent).
	ShapeColor	cdp.RGBA	`json:"shapeColor,omitempty"`

	// The shape margin fill color (default: transparent).
	ShapeMarginColor	cdp.RGBA	`json:"shapeMarginColor,omitempty"`

	// The grid layout color (default: transparent).
	CssGridColor	cdp.RGBA	`json:"cssGridColor,omitempty"`
}

// 
type InspectMode string
