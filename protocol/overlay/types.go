package overlay

import (

	"github.com/diiyw/gator/protocol/dom"

)

// Configuration data for the highlighting of page elements.
type HighlightConfig struct {
	
	// Whether the node info tooltip should be shown (default: false).
	
	ShowInfo	bool	`json:"showInfo"`
	
	// Whether the rulers should be shown (default: false).
	
	ShowRulers	bool	`json:"showRulers"`
	
	// Whether the extension lines from node to the rulers should be shown (default: false).
	
	ShowExtensionLines	bool	`json:"showExtensionLines"`
	
	
	
	DisplayAsMaterial	bool	`json:"displayAsMaterial"`
	
	// The content box highlight fill color (default: transparent).
	
	ContentColor	dom.RGBA	`json:"contentColor"`
	
	// The padding highlight fill color (default: transparent).
	
	PaddingColor	dom.RGBA	`json:"paddingColor"`
	
	// The border highlight fill color (default: transparent).
	
	BorderColor	dom.RGBA	`json:"borderColor"`
	
	// The margin highlight fill color (default: transparent).
	
	MarginColor	dom.RGBA	`json:"marginColor"`
	
	// The event target element highlight fill color (default: transparent).
	
	EventTargetColor	dom.RGBA	`json:"eventTargetColor"`
	
	// The shape outside fill color (default: transparent).
	
	ShapeColor	dom.RGBA	`json:"shapeColor"`
	
	// The shape margin fill color (default: transparent).
	
	ShapeMarginColor	dom.RGBA	`json:"shapeMarginColor"`
	
	// Selectors to highlight relevant nodes.
	
	SelectorList	string	`json:"selectorList"`
	
	// The grid layout color (default: transparent).
	
	CssGridColor	dom.RGBA	`json:"cssGridColor"`
	
}	

// 
type InspectMode string	

