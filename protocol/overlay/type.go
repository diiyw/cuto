package overlay
// Configuration data for the highlighting of page elements.
type HighlightConfig  struct {

	// Whether the node info tooltip should be shown (default: false).
	ShowInfo	bool	`json:"showInfo"`

	// Whether the node styles in the tooltip (default: false).
	ShowStyles	bool	`json:"showStyles"`

	// Whether the rulers should be shown (default: false).
	ShowRulers	bool	`json:"showRulers"`

	// Whether the extension lines from node to the rulers should be shown (default: false).
	ShowExtensionLines	bool	`json:"showExtensionLines"`

	// The content box highlight fill color (default: transparent).
	ContentColor	interface{}	`json:"contentColor"`

	// The padding highlight fill color (default: transparent).
	PaddingColor	interface{}	`json:"paddingColor"`

	// The border highlight fill color (default: transparent).
	BorderColor	interface{}	`json:"borderColor"`

	// The margin highlight fill color (default: transparent).
	MarginColor	interface{}	`json:"marginColor"`

	// The event target element highlight fill color (default: transparent).
	EventTargetColor	interface{}	`json:"eventTargetColor"`

	// The shape outside fill color (default: transparent).
	ShapeColor	interface{}	`json:"shapeColor"`

	// The shape margin fill color (default: transparent).
	ShapeMarginColor	interface{}	`json:"shapeMarginColor"`

	// The grid layout color (default: transparent).
	CssGridColor	interface{}	`json:"cssGridColor"`
}

// 
type InspectMode string
