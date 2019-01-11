package css

import (

	"github.com/diiyw/gator/protocol/dom"

	"github.com/diiyw/gator/protocol/frame"

)
const (
	
	// Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the
	// position specified by `location`.
	AddRule = "CSS.addRule"
	
	// Returns all class names from specified stylesheet.
	CollectClassNames = "CSS.collectClassNames"
	
	// Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
	CreateStyleSheet = "CSS.createStyleSheet"
	
	// Disables the CSS agent for the given page.
	Disable = "CSS.disable"
	
	// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been
	// enabled until the result of this command is received.
	Enable = "CSS.enable"
	
	// Ensures that the given node will have specified pseudo-classes whenever its style is computed by
	// the browser.
	ForcePseudoState = "CSS.forcePseudoState"
	
	
	GetBackgroundColors = "CSS.getBackgroundColors"
	
	// Returns the computed style for a DOM node identified by `nodeId`.
	GetComputedStyleForNode = "CSS.getComputedStyleForNode"
	
	// Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM
	// attributes) for a DOM node identified by `nodeId`.
	GetInlineStylesForNode = "CSS.getInlineStylesForNode"
	
	// Returns requested styles for a DOM node identified by `nodeId`.
	GetMatchedStylesForNode = "CSS.getMatchedStylesForNode"
	
	// Returns all media queries parsed by the rendering engine.
	GetMediaQueries = "CSS.getMediaQueries"
	
	// Requests information about platform fonts which we used to render child TextNodes in the given
	// node.
	GetPlatformFontsForNode = "CSS.getPlatformFontsForNode"
	
	// Returns the current textual content for a stylesheet.
	GetStyleSheetText = "CSS.getStyleSheetText"
	
	// Find a rule with the given active property for the given node and set the new value for this
	// property
	SetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"
	
	// Modifies the keyframe rule key text.
	SetKeyframeKey = "CSS.setKeyframeKey"
	
	// Modifies the rule selector.
	SetMediaText = "CSS.setMediaText"
	
	// Modifies the rule selector.
	SetRuleSelector = "CSS.setRuleSelector"
	
	// Sets the new stylesheet text.
	SetStyleSheetText = "CSS.setStyleSheetText"
	
	// Applies specified style edits one after another in the given order.
	SetStyleTexts = "CSS.setStyleTexts"
	
	// Enables the selector recording.
	StartRuleUsageTracking = "CSS.startRuleUsageTracking"
	
	// Stop tracking rule usage and return the list of rules that were used since last call to
	// `takeCoverageDelta` (or since start of coverage instrumentation)
	StopRuleUsageTracking = "CSS.stopRuleUsageTracking"
	
	// Obtain list of rules that became used since last call to this method (or since start of coverage
	// instrumentation)
	TakeCoverageDelta = "CSS.takeCoverageDelta"
	
)

// AddRule parameters
type AddRuleParams struct {
	
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
	// The text of a new rule.
	RuleText	string	`json:"ruleText"`
	
	// Text position of a new rule in the target style sheet.
	Location	SourceRange	`json:"location"`
	
}

// AddRule returns
type AddRuleReturns struct {
	
	// The newly created rule.
	Rule	CSSRule	`json:"rule"`
	
}

// CollectClassNames parameters
type CollectClassNamesParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
}

// CollectClassNames returns
type CollectClassNamesReturns struct {
	
	// Class name list.
	ClassNames	[]string	`json:"classNames"`
	
}

// CreateStyleSheet parameters
type CreateStyleSheetParams struct {
	
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId	frame.FrameId	`json:"frameId"`
	
}

// CreateStyleSheet returns
type CreateStyleSheetReturns struct {
	
	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// ForcePseudoState parameters
type ForcePseudoStateParams struct {
	
	// The element id for which to force the pseudo state.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses	[]string	`json:"forcedPseudoClasses"`
	
}

// ForcePseudoState returns
type ForcePseudoStateReturns struct {
	
}

// GetBackgroundColors parameters
type GetBackgroundColorsParams struct {
	
	// Id of the node to get background colors for.
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetBackgroundColors returns
type GetBackgroundColorsReturns struct {
	
	// The range of background colors behind this element, if it contains any visible text. If no
	// visible text is present, this will be undefined. In the case of a flat background color,
	// this will consist of simply that color. In the case of a gradient, this will consist of each
	// of the color stops. For anything more complicated, this will be an empty array. Images will
	// be ignored (as if the image had failed to load).
	BackgroundColors	[]string	`json:"backgroundColors"`
	
	// The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontSize	string	`json:"computedFontSize"`
	
	// The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or
	// '100').
	ComputedFontWeight	string	`json:"computedFontWeight"`
	
	// The computed font size for the document body, as a computed CSS value string (e.g. '16px').
	ComputedBodyFontSize	string	`json:"computedBodyFontSize"`
	
}

// GetComputedStyleForNode parameters
type GetComputedStyleForNodeParams struct {
	
	
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetComputedStyleForNode returns
type GetComputedStyleForNodeReturns struct {
	
	// Computed style for the specified DOM node.
	ComputedStyle	[]CSSComputedStyleProperty	`json:"computedStyle"`
	
}

// GetInlineStylesForNode parameters
type GetInlineStylesForNodeParams struct {
	
	
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetInlineStylesForNode returns
type GetInlineStylesForNodeReturns struct {
	
	// Inline style for the specified DOM node.
	InlineStyle	CSSStyle	`json:"inlineStyle"`
	
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle	CSSStyle	`json:"attributesStyle"`
	
}

// GetMatchedStylesForNode parameters
type GetMatchedStylesForNodeParams struct {
	
	
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetMatchedStylesForNode returns
type GetMatchedStylesForNodeReturns struct {
	
	// Inline style for the specified DOM node.
	InlineStyle	CSSStyle	`json:"inlineStyle"`
	
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle	CSSStyle	`json:"attributesStyle"`
	
	// CSS rules matching this node, from all applicable stylesheets.
	MatchedCSSRules	[]RuleMatch	`json:"matchedCSSRules"`
	
	// Pseudo style matches for this node.
	PseudoElements	[]PseudoElementMatches	`json:"pseudoElements"`
	
	// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	Inherited	[]InheritedStyleEntry	`json:"inherited"`
	
	// A list of CSS keyframed animations matching this node.
	CssKeyframesRules	[]CSSKeyframesRule	`json:"cssKeyframesRules"`
	
}

// GetMediaQueries parameters
type GetMediaQueriesParams struct {
	
}

// GetMediaQueries returns
type GetMediaQueriesReturns struct {
	
	
	Medias	[]CSSMedia	`json:"medias"`
	
}

// GetPlatformFontsForNode parameters
type GetPlatformFontsForNodeParams struct {
	
	
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetPlatformFontsForNode returns
type GetPlatformFontsForNodeReturns struct {
	
	// Usage statistics for every employed platform font.
	Fonts	[]PlatformFontUsage	`json:"fonts"`
	
}

// GetStyleSheetText parameters
type GetStyleSheetTextParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
}

// GetStyleSheetText returns
type GetStyleSheetTextReturns struct {
	
	// The stylesheet text.
	Text	string	`json:"text"`
	
}

// SetEffectivePropertyValueForNode parameters
type SetEffectivePropertyValueForNodeParams struct {
	
	// The element id for which to set property.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	
	PropertyName	string	`json:"propertyName"`
	
	
	Value	string	`json:"value"`
	
}

// SetEffectivePropertyValueForNode returns
type SetEffectivePropertyValueForNodeReturns struct {
	
}

// SetKeyframeKey parameters
type SetKeyframeKeyParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
	
	Range	SourceRange	`json:"range"`
	
	
	KeyText	string	`json:"keyText"`
	
}

// SetKeyframeKey returns
type SetKeyframeKeyReturns struct {
	
	// The resulting key text after modification.
	KeyText	Value	`json:"keyText"`
	
}

// SetMediaText parameters
type SetMediaTextParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
	
	Range	SourceRange	`json:"range"`
	
	
	Text	string	`json:"text"`
	
}

// SetMediaText returns
type SetMediaTextReturns struct {
	
	// The resulting CSS media rule after modification.
	Media	CSSMedia	`json:"media"`
	
}

// SetRuleSelector parameters
type SetRuleSelectorParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
	
	Range	SourceRange	`json:"range"`
	
	
	Selector	string	`json:"selector"`
	
}

// SetRuleSelector returns
type SetRuleSelectorReturns struct {
	
	// The resulting selector list after modification.
	SelectorList	SelectorList	`json:"selectorList"`
	
}

// SetStyleSheetText parameters
type SetStyleSheetTextParams struct {
	
	
	StyleSheetId	StyleSheetId	`json:"styleSheetId"`
	
	
	Text	string	`json:"text"`
	
}

// SetStyleSheetText returns
type SetStyleSheetTextReturns struct {
	
	// URL of source map associated with script (if any).
	SourceMapURL	string	`json:"sourceMapURL"`
	
}

// SetStyleTexts parameters
type SetStyleTextsParams struct {
	
	
	Edits	[]StyleDeclarationEdit	`json:"edits"`
	
}

// SetStyleTexts returns
type SetStyleTextsReturns struct {
	
	// The resulting styles after modification.
	Styles	[]CSSStyle	`json:"styles"`
	
}

// StartRuleUsageTracking parameters
type StartRuleUsageTrackingParams struct {
	
}

// StartRuleUsageTracking returns
type StartRuleUsageTrackingReturns struct {
	
}

// StopRuleUsageTracking parameters
type StopRuleUsageTrackingParams struct {
	
}

// StopRuleUsageTracking returns
type StopRuleUsageTrackingReturns struct {
	
	
	RuleUsage	[]RuleUsage	`json:"ruleUsage"`
	
}

// TakeCoverageDelta parameters
type TakeCoverageDeltaParams struct {
	
}

// TakeCoverageDelta returns
type TakeCoverageDeltaReturns struct {
	
	
	Coverage	[]RuleUsage	`json:"coverage"`
	
}

