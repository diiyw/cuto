package domsnapshot

import (
	"github.com/diiyw/cuto/protocol/domdebugger"
	"github.com/diiyw/cuto/protocol/dom"
)

// A Node in the DOM tree.
type DOMNode  struct {

	// `Node`'s nodeType.
	NodeType	int	`json:"nodeType"`

	// `Node`'s nodeName.
	NodeName	string	`json:"nodeName"`

	// `Node`'s nodeValue.
	NodeValue	string	`json:"nodeValue"`

	// Only set for textarea elements, contains the text value.
	TextValue	string	`json:"textValue"`

	// Only set for input elements, contains the input's associated text value.
	InputValue	string	`json:"inputValue"`

	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked	bool	`json:"inputChecked"`

	// Only set for option elements, indicates if the element has been selected
	OptionSelected	bool	`json:"optionSelected"`

	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeId	interface{}	`json:"backendNodeId"`

	// The indexes of the node's child nodes in the `domNodes` array returned by `getSnapshot`, if
	// any.
	ChildNodeIndexes	[]int	`json:"childNodeIndexes"`

	// Attributes of an `Element` node.
	Attributes	[]*NameValue	`json:"attributes"`

	// Indexes of pseudo elements associated with this node in the `domNodes` array returned by
	// `getSnapshot`, if any.
	PseudoElementIndexes	[]int	`json:"pseudoElementIndexes"`

	// The index of the node's related layout tree node in the `layoutTreeNodes` array returned by
	// `getSnapshot`, if any.
	LayoutNodeIndex	int	`json:"layoutNodeIndex"`

	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL	string	`json:"documentURL"`

	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL	string	`json:"baseURL"`

	// Only set for documents, contains the document's content language.
	ContentLanguage	string	`json:"contentLanguage"`

	// Only set for documents, contains the document's character set encoding.
	DocumentEncoding	string	`json:"documentEncoding"`

	// `DocumentType` node's publicId.
	PublicId	string	`json:"publicId"`

	// `DocumentType` node's systemId.
	SystemId	string	`json:"systemId"`

	// Frame ID for frame owner elements and also for the document node.
	FrameId	interface{}	`json:"frameId"`

	// The index of a frame owner element's content document in the `domNodes` array returned by
	// `getSnapshot`, if any.
	ContentDocumentIndex	int	`json:"contentDocumentIndex"`

	// Type of a pseudo element node.
	PseudoType	interface{}	`json:"pseudoType"`

	// Shadow root type.
	ShadowRootType	interface{}	`json:"shadowRootType"`

	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable	bool	`json:"isClickable"`

	// Details of the node's event listeners, if any.
	EventListeners	[]*domdebugger.EventListener	`json:"eventListeners"`

	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL	string	`json:"currentSourceURL"`

	// The url of the script (if any) that generates this node.
	OriginURL	string	`json:"originURL"`

	// Scroll offsets, set when this node is a Document.
	ScrollOffsetX	float64	`json:"scrollOffsetX"`

	// 
	ScrollOffsetY	float64	`json:"scrollOffsetY"`
}

// Details of post layout rendered text positions. The exact layout should not be regarded as
	// stable and may change between versions.
type InlineTextBox  struct {

	// The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox	interface{}	`json:"boundingBox"`

	// The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	StartCharacterIndex	int	`json:"startCharacterIndex"`

	// The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters	int	`json:"numCharacters"`
}

// Details of an element in the DOM tree with a LayoutObject.
type LayoutTreeNode  struct {

	// The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	DomNodeIndex	int	`json:"domNodeIndex"`

	// The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	BoundingBox	interface{}	`json:"boundingBox"`

	// Contents of the LayoutText, if any.
	LayoutText	string	`json:"layoutText"`

	// The post-layout inline text nodes, if any.
	InlineTextNodes	[]*InlineTextBox	`json:"inlineTextNodes"`

	// Index into the `computedStyles` array returned by `getSnapshot`.
	StyleIndex	int	`json:"styleIndex"`

	// Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// getSnapshot was true.
	PaintOrder	int	`json:"paintOrder"`

	// Set to true to indicate the element begins a new stacking context.
	IsStackingContext	bool	`json:"isStackingContext"`
}

// A subset of the full ComputedStyle as defined by the request whitelist.
type ComputedStyle  struct {

	// Name/value pairs of computed style properties.
	Properties	[]*NameValue	`json:"properties"`
}

// A name/value pair.
type NameValue  struct {

	// Attribute/property name.
	Name	string	`json:"name"`

	// Attribute/property value.
	Value	string	`json:"value"`
}

// Index of the string in the strings table.
type StringIndex int

// Index of the string in the strings table.
type ArrayOfStrings 	[]*StringIndex

// Data that is only present on rare nodes.
type RareStringData  struct {

	// 
	Index	[]int	`json:"index"`

	// 
	Value	[]*StringIndex	`json:"value"`
}

// 
type RareBooleanData  struct {

	// 
	Index	[]int	`json:"index"`
}

// 
type RareIntegerData  struct {

	// 
	Index	[]int	`json:"index"`

	// 
	Value	[]int	`json:"value"`
}

// 
type Rectangle 	[]float64

// Document snapshot.
type DocumentSnapshot  struct {

	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL	StringIndex	`json:"documentURL"`

	// Document title.
	Title	StringIndex	`json:"title"`

	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL	StringIndex	`json:"baseURL"`

	// Contains the document's content language.
	ContentLanguage	StringIndex	`json:"contentLanguage"`

	// Contains the document's character set encoding.
	EncodingName	StringIndex	`json:"encodingName"`

	// `DocumentType` node's publicId.
	PublicId	StringIndex	`json:"publicId"`

	// `DocumentType` node's systemId.
	SystemId	StringIndex	`json:"systemId"`

	// Frame ID for frame owner elements and also for the document node.
	FrameId	StringIndex	`json:"frameId"`

	// A table with dom nodes.
	Nodes	NodeTreeSnapshot	`json:"nodes"`

	// The nodes in the layout tree.
	Layout	LayoutTreeSnapshot	`json:"layout"`

	// The post-layout inline text nodes.
	TextBoxes	TextBoxSnapshot	`json:"textBoxes"`

	// Horizontal scroll offset.
	ScrollOffsetX	float64	`json:"scrollOffsetX"`

	// Vertical scroll offset.
	ScrollOffsetY	float64	`json:"scrollOffsetY"`

	// Document content width.
	ContentWidth	float64	`json:"contentWidth"`

	// Document content height.
	ContentHeight	float64	`json:"contentHeight"`
}

// Table containing nodes.
type NodeTreeSnapshot  struct {

	// Parent node index.
	ParentIndex	[]int	`json:"parentIndex"`

	// `Node`'s nodeType.
	NodeType	[]int	`json:"nodeType"`

	// `Node`'s nodeName.
	NodeName	[]*StringIndex	`json:"nodeName"`

	// `Node`'s nodeValue.
	NodeValue	[]*StringIndex	`json:"nodeValue"`

	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeId	[]*dom.BackendNodeId	`json:"backendNodeId"`

	// Attributes of an `Element` node. Flatten name, value pairs.
	Attributes	[]*ArrayOfStrings	`json:"attributes"`

	// Only set for textarea elements, contains the text value.
	TextValue	RareStringData	`json:"textValue"`

	// Only set for input elements, contains the input's associated text value.
	InputValue	RareStringData	`json:"inputValue"`

	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked	RareBooleanData	`json:"inputChecked"`

	// Only set for option elements, indicates if the element has been selected
	OptionSelected	RareBooleanData	`json:"optionSelected"`

	// The index of the document in the list of the snapshot documents.
	ContentDocumentIndex	RareIntegerData	`json:"contentDocumentIndex"`

	// Type of a pseudo element node.
	PseudoType	RareStringData	`json:"pseudoType"`

	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click
	// event listeners attached via JavaScript as well as anchor tags that naturally navigate when
	// clicked.
	IsClickable	RareBooleanData	`json:"isClickable"`

	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL	RareStringData	`json:"currentSourceURL"`

	// The url of the script (if any) that generates this node.
	OriginURL	RareStringData	`json:"originURL"`
}

// Table of details of an element in the DOM tree with a LayoutObject.
type LayoutTreeSnapshot  struct {

	// Index of the corresponding node in the `NodeTreeSnapshot` array returned by `captureSnapshot`.
	NodeIndex	[]int	`json:"nodeIndex"`

	// Array of indexes specifying computed style strings, filtered according to the `computedStyles` parameter passed to `captureSnapshot`.
	Styles	[]*ArrayOfStrings	`json:"styles"`

	// The absolute position bounding box.
	Bounds	[]*Rectangle	`json:"bounds"`

	// Contents of the LayoutText, if any.
	Text	[]*StringIndex	`json:"text"`

	// Stacking context information.
	StackingContexts	RareBooleanData	`json:"stackingContexts"`

	// Global paint order index, which is determined by the stacking order of the nodes. Nodes
	// that are painted together will have the same index. Only provided if includePaintOrder in
	// captureSnapshot was true.
	PaintOrders	[]int	`json:"paintOrders"`

	// The offset rect of nodes. Only available when includeDOMRects is set to true
	OffsetRects	[]*Rectangle	`json:"offsetRects"`

	// The scroll rect of nodes. Only available when includeDOMRects is set to true
	ScrollRects	[]*Rectangle	`json:"scrollRects"`

	// The client rect of nodes. Only available when includeDOMRects is set to true
	ClientRects	[]*Rectangle	`json:"clientRects"`
}

// Table of details of the post layout rendered text positions. The exact layout should not be regarded as
	// stable and may change between versions.
type TextBoxSnapshot  struct {

	// Index of the layout tree node that owns this box collection.
	LayoutIndex	[]int	`json:"layoutIndex"`

	// The absolute position bounding box.
	Bounds	[]*Rectangle	`json:"bounds"`

	// The starting index in characters, for this post layout textbox substring. Characters that
	// would be represented as a surrogate pair in UTF-16 have length 2.
	Start	[]int	`json:"start"`

	// The number of characters in this post layout textbox substring. Characters that would be
	// represented as a surrogate pair in UTF-16 have length 2.
	Length	[]int	`json:"length"`
}
