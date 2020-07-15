package css
// 
type StyleSheetId string

// Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent
	// stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via
	// inspector" rules), "regular" for regular stylesheets.
type StyleSheetOrigin string

// CSS rule collection for a single pseudo style.
type PseudoElementMatches  struct {

	// Pseudo element type.
	PseudoType	interface{}	`json:"pseudoType"`

	// Matches of CSS rules applicable to the pseudo style.
	Matches	[]*RuleMatch	`json:"matches"`
}

// Inherited CSS rule collection from ancestor node.
type InheritedStyleEntry  struct {

	// The ancestor node's inline style, if any, in the style inheritance chain.
	InlineStyle	CSSStyle	`json:"inlineStyle"`

	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules	[]*RuleMatch	`json:"matchedCSSRules"`
}

// Match data for a CSS rule.
type RuleMatch  struct {

	// CSS rule in the match.
	Rule	CSSRule	`json:"rule"`

	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors	[]int	`json:"matchingSelectors"`
}

// Data for a simple selector (these are delimited by commas in a selector list).
type Value  struct {

	// Value text.
	Text	string	`json:"text"`

	// Value range in the underlying resource (if available).
	Range	SourceRange	`json:"range"`
}

// Selector list data.
type SelectorList  struct {

	// Selectors in the list.
	Selectors	[]*Value	`json:"selectors"`

	// Rule selector text.
	Text	string	`json:"text"`
}

// CSS stylesheet metainformation.
type CSSStyleSheetHeader  struct {

	// The stylesheet identifier.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// Owner frame identifier.
	FrameId	interface{}	`json:"frameId"`

	// Stylesheet resource URL.
	SourceURL	string	`json:"sourceURL"`

	// URL of source map associated with the stylesheet (if any).
	SourceMapURL	string	`json:"sourceMapURL"`

	// Stylesheet origin.
	Origin	StyleSheetOrigin	`json:"origin"`

	// Stylesheet title.
	Title	string	`json:"title"`

	// The backend id for the owner node of the stylesheet.
	OwnerNode	interface{}	`json:"ownerNode"`

	// Denotes whether the stylesheet is disabled.
	Disabled	bool	`json:"disabled"`

	// Whether the sourceURL field value comes from the sourceURL comment.
	HasSourceURL	bool	`json:"hasSourceURL"`

	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for
	// document.written STYLE tags.
	IsInline	bool	`json:"isInline"`

	// Line offset of the stylesheet within the resource (zero based).
	StartLine	float64	`json:"startLine"`

	// Column offset of the stylesheet within the resource (zero based).
	StartColumn	float64	`json:"startColumn"`

	// Size of the content (in characters).
	Length	float64	`json:"length"`

	// Line offset of the end of the stylesheet within the resource (zero based).
	EndLine	float64	`json:"endLine"`

	// Column offset of the end of the stylesheet within the resource (zero based).
	EndColumn	float64	`json:"endColumn"`
}

// CSS rule representation.
type CSSRule  struct {

	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// Rule selector data.
	SelectorList	SelectorList	`json:"selectorList"`

	// Parent stylesheet's origin.
	Origin	StyleSheetOrigin	`json:"origin"`

	// Associated style declaration.
	Style	CSSStyle	`json:"style"`

	// Media list array (for rules involving media queries). The array enumerates media queries
	// starting with the innermost one, going outwards.
	Media	[]*CSSMedia	`json:"media"`
}

// CSS coverage information.
type RuleUsage  struct {

	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset	float64	`json:"startOffset"`

	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset	float64	`json:"endOffset"`

	// Indicates whether the rule was actually used by some element in the page.
	Used	bool	`json:"used"`
}

// Text range within a resource. All numbers are zero-based.
type SourceRange  struct {

	// Start line of range.
	StartLine	int	`json:"startLine"`

	// Start column of range (inclusive).
	StartColumn	int	`json:"startColumn"`

	// End line of range
	EndLine	int	`json:"endLine"`

	// End column of range (exclusive).
	EndColumn	int	`json:"endColumn"`
}

// 
type ShorthandEntry  struct {

	// Shorthand name.
	Name	string	`json:"name"`

	// Shorthand value.
	Value	string	`json:"value"`

	// Whether the property has "!important" annotation (implies `false` if absent).
	Important	bool	`json:"important"`
}

// 
type CSSComputedStyleProperty  struct {

	// Computed style property name.
	Name	string	`json:"name"`

	// Computed style property value.
	Value	string	`json:"value"`
}

// CSS style representation.
type CSSStyle  struct {

	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// CSS properties in the style.
	CssProperties	[]*CSSProperty	`json:"cssProperties"`

	// Computed values for all shorthands found in the style.
	ShorthandEntries	[]*ShorthandEntry	`json:"shorthandEntries"`

	// Style declaration text (if available).
	CssText	string	`json:"cssText"`

	// Style declaration range in the enclosing stylesheet (if available).
	Range	SourceRange	`json:"range"`
}

// CSS property declaration data.
type CSSProperty  struct {

	// The property name.
	Name	string	`json:"name"`

	// The property value.
	Value	string	`json:"value"`

	// Whether the property has "!important" annotation (implies `false` if absent).
	Important	bool	`json:"important"`

	// Whether the property is implicit (implies `false` if absent).
	Implicit	bool	`json:"implicit"`

	// The full property text as specified in the style.
	Text	string	`json:"text"`

	// Whether the property is understood by the browser (implies `true` if absent).
	ParsedOk	bool	`json:"parsedOk"`

	// Whether the property is disabled by the user (present for source-based properties only).
	Disabled	bool	`json:"disabled"`

	// The entire property range in the enclosing style declaration (if available).
	Range	SourceRange	`json:"range"`
}

// CSS media rule descriptor.
type CSSMedia  struct {

	// Media query text.
	Text	string	`json:"text"`

	// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if
	// specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked
	// stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline
	// stylesheet's STYLE tag.
	Source	string	`json:"source"`

	// URL of the document containing the media query description.
	SourceURL	string	`json:"sourceURL"`

	// The associated rule (@media or @import) header range in the enclosing stylesheet (if
	// available).
	Range	SourceRange	`json:"range"`

	// Identifier of the stylesheet containing this object (if exists).
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// Array of media queries.
	MediaList	[]*MediaQuery	`json:"mediaList"`
}

// Media query descriptor.
type MediaQuery  struct {

	// Array of media query expressions.
	Expressions	[]*MediaQueryExpression	`json:"expressions"`

	// Whether the media query condition is satisfied.
	Active	bool	`json:"active"`
}

// Media query expression descriptor.
type MediaQueryExpression  struct {

	// Media query expression value.
	Value	float64	`json:"value"`

	// Media query expression units.
	Unit	string	`json:"unit"`

	// Media query expression feature.
	Feature	string	`json:"feature"`

	// The associated range of the value text in the enclosing stylesheet (if available).
	ValueRange	SourceRange	`json:"valueRange"`

	// Computed length of media query expression (if applicable).
	ComputedLength	float64	`json:"computedLength"`
}

// Information about amount of glyphs that were rendered with given font.
type PlatformFontUsage  struct {

	// FontFamily's family name reported by platform.
	FamilyName	string	`json:"familyName"`

	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont	bool	`json:"isCustomFont"`

	// Amount of glyphs that were rendered with this font.
	GlyphCount	float64	`json:"glyphCount"`
}

// Properties of a web font: https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions
type FontFace  struct {

	// The font-family.
	FontFamily	string	`json:"fontFamily"`

	// The font-style.
	FontStyle	string	`json:"fontStyle"`

	// The font-variant.
	FontVariant	string	`json:"fontVariant"`

	// The font-weight.
	FontWeight	string	`json:"fontWeight"`

	// The font-stretch.
	FontStretch	string	`json:"fontStretch"`

	// The unicode-range.
	UnicodeRange	string	`json:"unicodeRange"`

	// The src.
	Src	string	`json:"src"`

	// The resolved platform font family
	PlatformFontFamily	string	`json:"platformFontFamily"`
}

// CSS keyframes rule representation.
type CSSKeyframesRule  struct {

	// Animation name.
	AnimationName	Value	`json:"animationName"`

	// List of keyframes.
	Keyframes	[]*CSSKeyframeRule	`json:"keyframes"`
}

// CSS keyframe rule representation.
type CSSKeyframeRule  struct {

	// The css style sheet identifier (absent for user agent stylesheet and user-specified
	// stylesheet rules) this rule came from.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// Parent stylesheet's origin.
	Origin	StyleSheetOrigin	`json:"origin"`

	// Associated key text.
	KeyText	Value	`json:"keyText"`

	// Associated style declaration.
	Style	CSSStyle	`json:"style"`
}

// A descriptor of operation to mutate style declaration text.
type StyleDeclarationEdit  struct {

	// The css style sheet identifier.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`

	// The range of the style text in the enclosing stylesheet.
	Range	SourceRange	`json:"range"`

	// New style text.
	Text	string	`json:"text"`
}
