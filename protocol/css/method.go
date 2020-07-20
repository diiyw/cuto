package css

import (
	"github.com/diiyw/cuto/protocol/cdp"
	"github.com/diiyw/cuto/protocol/dom"
)


// Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the
// position specified by `location`.
const AddRule = "CSS.addRule"

type AddRuleParams struct {

	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`

	// The text of a new rule.
	RuleText 	string	`json:"ruleText"`

	// Text position of a new rule in the target style sheet.
	Location 	SourceRange	`json:"location"`
}

type AddRuleResult struct {

	// The newly created rule.
	Rule 	CSSRule	`json:"rule"`
}

// Returns all class names from specified stylesheet.
const CollectClassNames = "CSS.collectClassNames"

type CollectClassNamesParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`
}

type CollectClassNamesResult struct {

	// Class name list.
	ClassNames 	[]string	`json:"classNames"`
}

// Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
const CreateStyleSheet = "CSS.createStyleSheet"

type CreateStyleSheetParams struct {

	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId 	cdp.FrameId	`json:"frameId"`
}

type CreateStyleSheetResult struct {

	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`
}

// Disables the CSS agent for the given page.
const Disable = "CSS.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been
// enabled until the result of this command is received.
const Enable = "CSS.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Ensures that the given node will have specified pseudo-classes whenever its style is computed by
// the browser.
const ForcePseudoState = "CSS.forcePseudoState"

type ForcePseudoStateParams struct {

	// The element id for which to force the pseudo state.
	NodeId 	dom.NodeId	`json:"nodeId"`

	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses 	[]string	`json:"forcedPseudoClasses"`
}

type ForcePseudoStateResult struct {

}

// 
const GetBackgroundColors = "CSS.getBackgroundColors"

type GetBackgroundColorsParams struct {

	// Id of the node to get background colors for.
	NodeId 	dom.NodeId	`json:"nodeId"`
}

type GetBackgroundColorsResult struct {

	// The range of background colors behind this element, if it contains any visible text. If no
	// visible text is present, this will be undefined. In the case of a flat background color,
	// this will consist of simply that color. In the case of a gradient, this will consist of each
	// of the color stops. For anything more complicated, this will be an empty array. Images will
	// be ignored (as if the image had failed to load).
	BackgroundColors 	[]string	`json:"backgroundColors"`
	// The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontSize 	string	`json:"computedFontSize"`
	// The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or
	// '100').
	ComputedFontWeight 	string	`json:"computedFontWeight"`
}

// Returns the computed style for a DOM node identified by `nodeId`.
const GetComputedStyleForNode = "CSS.getComputedStyleForNode"

type GetComputedStyleForNodeParams struct {

	// 
	NodeId 	dom.NodeId	`json:"nodeId"`
}

type GetComputedStyleForNodeResult struct {

	// Computed style for the specified DOM node.
	ComputedStyle 	[]*CSSComputedStyleProperty	`json:"computedStyle"`
}

// Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM
// attributes) for a DOM node identified by `nodeId`.
const GetInlineStylesForNode = "CSS.getInlineStylesForNode"

type GetInlineStylesForNodeParams struct {

	// 
	NodeId 	dom.NodeId	`json:"nodeId"`
}

type GetInlineStylesForNodeResult struct {

	// Inline style for the specified DOM node.
	InlineStyle 	CSSStyle	`json:"inlineStyle"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle 	CSSStyle	`json:"attributesStyle"`
}

// Returns requested styles for a DOM node identified by `nodeId`.
const GetMatchedStylesForNode = "CSS.getMatchedStylesForNode"

type GetMatchedStylesForNodeParams struct {

	// 
	NodeId 	dom.NodeId	`json:"nodeId"`
}

type GetMatchedStylesForNodeResult struct {

	// Inline style for the specified DOM node.
	InlineStyle 	CSSStyle	`json:"inlineStyle"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle 	CSSStyle	`json:"attributesStyle"`
	// CSS rules matching this node, from all applicable stylesheets.
	MatchedCSSRules 	[]*RuleMatch	`json:"matchedCSSRules"`
	// Pseudo style matches for this node.
	PseudoElements 	[]*PseudoElementMatches	`json:"pseudoElements"`
	// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	Inherited 	[]*InheritedStyleEntry	`json:"inherited"`
	// A list of CSS keyframed animations matching this node.
	CssKeyframesRules 	[]*CSSKeyframesRule	`json:"cssKeyframesRules"`
}

// Returns all media queries parsed by the rendering engine.
const GetMediaQueries = "CSS.getMediaQueries"

type GetMediaQueriesParams struct {
}

type GetMediaQueriesResult struct {

	// 
	Medias 	[]*CSSMedia	`json:"medias"`
}

// Requests information about platform fonts which we used to render child TextNodes in the given
// node.
const GetPlatformFontsForNode = "CSS.getPlatformFontsForNode"

type GetPlatformFontsForNodeParams struct {

	// 
	NodeId 	dom.NodeId	`json:"nodeId"`
}

type GetPlatformFontsForNodeResult struct {

	// Usage statistics for every employed platform font.
	Fonts 	[]*PlatformFontUsage	`json:"fonts"`
}

// Returns the current textual content for a stylesheet.
const GetStyleSheetText = "CSS.getStyleSheetText"

type GetStyleSheetTextParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`
}

type GetStyleSheetTextResult struct {

	// The stylesheet text.
	Text 	string	`json:"text"`
}

// Find a rule with the given active property for the given node and set the new value for this
// property
const SetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"

type SetEffectivePropertyValueForNodeParams struct {

	// The element id for which to set property.
	NodeId 	dom.NodeId	`json:"nodeId"`

	// 
	PropertyName 	string	`json:"propertyName"`

	// 
	Value 	string	`json:"value"`
}

type SetEffectivePropertyValueForNodeResult struct {

}

// Modifies the keyframe rule key text.
const SetKeyframeKey = "CSS.setKeyframeKey"

type SetKeyframeKeyParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`

	// 
	Range 	SourceRange	`json:"range"`

	// 
	KeyText 	string	`json:"keyText"`
}

type SetKeyframeKeyResult struct {

	// The resulting key text after modification.
	KeyText 	Value	`json:"keyText"`
}

// Modifies the rule selector.
const SetMediaText = "CSS.setMediaText"

type SetMediaTextParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`

	// 
	Range 	SourceRange	`json:"range"`

	// 
	Text 	string	`json:"text"`
}

type SetMediaTextResult struct {

	// The resulting CSS media rule after modification.
	Media 	CSSMedia	`json:"media"`
}

// Modifies the rule selector.
const SetRuleSelector = "CSS.setRuleSelector"

type SetRuleSelectorParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`

	// 
	Range 	SourceRange	`json:"range"`

	// 
	Selector 	string	`json:"selector"`
}

type SetRuleSelectorResult struct {

	// The resulting selector list after modification.
	SelectorList 	SelectorList	`json:"selectorList"`
}

// Sets the new stylesheet text.
const SetStyleSheetText = "CSS.setStyleSheetText"

type SetStyleSheetTextParams struct {

	// 
	StyleSheetId 	StyleSheetId	`json:"styleSheetId"`

	// 
	Text 	string	`json:"text"`
}

type SetStyleSheetTextResult struct {

	// URL of source map associated with script (if any).
	SourceMapURL 	string	`json:"sourceMapURL"`
}

// Applies specified style edits one after another in the given order.
const SetStyleTexts = "CSS.setStyleTexts"

type SetStyleTextsParams struct {

	// 
	Edits 	[]*StyleDeclarationEdit	`json:"edits"`
}

type SetStyleTextsResult struct {

	// The resulting styles after modification.
	Styles 	[]*CSSStyle	`json:"styles"`
}

// Enables the selector recording.
const StartRuleUsageTracking = "CSS.startRuleUsageTracking"

type StartRuleUsageTrackingParams struct {
}

type StartRuleUsageTrackingResult struct {

}

// Stop tracking rule usage and return the list of rules that were used since last call to
// `takeCoverageDelta` (or since start of coverage instrumentation)
const StopRuleUsageTracking = "CSS.stopRuleUsageTracking"

type StopRuleUsageTrackingParams struct {
}

type StopRuleUsageTrackingResult struct {

	// 
	RuleUsage 	[]*RuleUsage	`json:"ruleUsage"`
}

// Obtain list of rules that became used since last call to this method (or since start of coverage
// instrumentation)
const TakeCoverageDelta = "CSS.takeCoverageDelta"

type TakeCoverageDeltaParams struct {
}

type TakeCoverageDeltaResult struct {

	// 
	Coverage 	[]*RuleUsage	`json:"coverage"`
}