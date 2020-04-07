package css

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded
// web font
const FontsUpdatedEvent = "CSS.fontsUpdated"
type FontsUpdatedParams struct {

	// The web font that has loaded.
	Font 	FontFace}



// Fires whenever a MediaQuery result changes (for example, after a browser window has been
// resized.) The current implementation considers only viewport-dependent media features.
const MediaQueryResultChangedEvent = "CSS.mediaQueryResultChanged"
type MediaQueryResultChangedParams struct {
}



// Fired whenever an active document stylesheet is added.
const StyleSheetAddedEvent = "CSS.styleSheetAdded"
type StyleSheetAddedParams struct {

	// Added stylesheet metainfo.
	Header 	CSSStyleSheetHeader}



// Fired whenever a stylesheet is changed as a result of the client operation.
const StyleSheetChangedEvent = "CSS.styleSheetChanged"
type StyleSheetChangedParams struct {

	// 
	StyleSheetId 	StyleSheetId}



// Fired whenever an active document stylesheet is removed.
const StyleSheetRemovedEvent = "CSS.styleSheetRemoved"
type StyleSheetRemovedParams struct {

	// Identifier of the removed stylesheet.
	StyleSheetId 	StyleSheetId}

