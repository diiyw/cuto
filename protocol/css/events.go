package css

const (
	
	// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded
	// web font
	FontsUpdatedEvent = "CSS.fontsUpdated"
	
	// Fires whenever a MediaQuery result changes (for example, after a browser window has been
	// resized.) The current implementation considers only viewport-dependent media features.
	MediaQueryResultChangedEvent = "CSS.mediaQueryResultChanged"
	
	// Fired whenever an active document stylesheet is added.
	StyleSheetAddedEvent = "CSS.styleSheetAdded"
	
	// Fired whenever a stylesheet is changed as a result of the client operation.
	StyleSheetChangedEvent = "CSS.styleSheetChanged"
	
	// Fired whenever an active document stylesheet is removed.
	StyleSheetRemovedEvent = "CSS.styleSheetRemoved"
	
)

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded
	// web font
type FontsUpdatedParams struct {
	
	// The web font that has loaded.
	Font	FontFace	`json:"font"`
	
}

// Fires whenever a MediaQuery result changes (for example, after a browser window has been
	// resized.) The current implementation considers only viewport-dependent media features.
type MediaQueryResultChangedParams struct {
	
}

// Fired whenever an active document stylesheet is added.
type StyleSheetAddedParams struct {
	
	// Added stylesheet metainfo.
	Header	CSSStyleSheetHeader	`json:"header"`
	
}

// Fired whenever a stylesheet is changed as a result of the client operation.
type StyleSheetChangedParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
}

// Fired whenever an active document stylesheet is removed.
type StyleSheetRemovedParams struct {
	
	// Identifier of the removed stylesheet.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
}

