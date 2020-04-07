package dom
// Unique DOM node identifier.
type NodeId int

// Unique DOM node identifier used to reference a node that may not have been pushed to the
	// front-end.
type BackendNodeId int

// Backend node with a friendly name.
type BackendNode  struct {

	// `Node`'s nodeType.
	NodeType	int	`json:"nodeType"`

	// `Node`'s nodeName.
	NodeName	string	`json:"nodeName"`

	// 
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
}

// Pseudo element type.
type PseudoType string

// Shadow root type.
type ShadowRootType string

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes.
	// DOMNode is a base node mirror type.
type Node  struct {

	// Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend
	// will only push node with given `id` once. It is aware of all requested nodes and will only
	// fire DOM events for nodes known to the client.
	NodeId	NodeId	`json:"nodeId"`

	// The id of the parent node if any.
	ParentId	NodeId	`json:"parentId"`

	// The BackendNodeId for this node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`

	// `Node`'s nodeType.
	NodeType	int	`json:"nodeType"`

	// `Node`'s nodeName.
	NodeName	string	`json:"nodeName"`

	// `Node`'s localName.
	LocalName	string	`json:"localName"`

	// `Node`'s nodeValue.
	NodeValue	string	`json:"nodeValue"`

	// Child count for `Container` nodes.
	ChildNodeCount	int	`json:"childNodeCount"`

	// Child nodes of this node when requested with children.
	Children	[]*Node	`json:"children"`

	// Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	Attributes	[]string	`json:"attributes"`

	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL	string	`json:"documentURL"`

	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL	string	`json:"baseURL"`

	// `DocumentType`'s publicId.
	PublicId	string	`json:"publicId"`

	// `DocumentType`'s systemId.
	SystemId	string	`json:"systemId"`

	// `DocumentType`'s internalSubset.
	InternalSubset	string	`json:"internalSubset"`

	// `Document`'s XML version in case of XML documents.
	XmlVersion	string	`json:"xmlVersion"`

	// `Attr`'s name.
	Name	string	`json:"name"`

	// `Attr`'s value.
	Value	string	`json:"value"`

	// Pseudo element type for this node.
	PseudoType	PseudoType	`json:"pseudoType"`

	// Shadow root type.
	ShadowRootType	ShadowRootType	`json:"shadowRootType"`

	// Frame ID for frame owner elements.
	FrameId	interface{}	`json:"frameId"`

	// Content document for frame owner elements.
	ContentDocument	*Node	`json:"contentDocument"`

	// Shadow root list for given element host.
	ShadowRoots	[]*Node	`json:"shadowRoots"`

	// Content document fragment for template elements.
	TemplateContent	*Node	`json:"templateContent"`

	// Pseudo elements associated with this node.
	PseudoElements	[]*Node	`json:"pseudoElements"`

	// Import document for the HTMLImport links.
	ImportedDocument	*Node	`json:"importedDocument"`

	// Distributed nodes for given insertion point.
	DistributedNodes	[]*BackendNode	`json:"distributedNodes"`

	// Whether the node is SVG.
	IsSVG	bool	`json:"isSVG"`
}

// A structure holding an RGBA color.
type RGBA  struct {

	// The red component, in the [0-255] range.
	R	int	`json:"r"`

	// The green component, in the [0-255] range.
	G	int	`json:"g"`

	// The blue component, in the [0-255] range.
	B	int	`json:"b"`

	// The alpha component, in the [0-1] range (default: 1).
	A	float64	`json:"a"`
}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise.
type Quad 	[]float64

// Box model.
type BoxModel  struct {

	// Content box
	Content	Quad	`json:"content"`

	// Padding box
	Padding	Quad	`json:"padding"`

	// Border box
	Border	Quad	`json:"border"`

	// Margin box
	Margin	Quad	`json:"margin"`

	// Node width
	Width	int	`json:"width"`

	// Node height
	Height	int	`json:"height"`

	// Shape outside coordinates
	ShapeOutside	ShapeOutsideInfo	`json:"shapeOutside"`
}

// CSS Shape Outside details.
type ShapeOutsideInfo  struct {

	// Shape bounds
	Bounds	Quad	`json:"bounds"`

	// Shape coordinate details
	Shape	[]interface{}	`json:"shape"`

	// Margin shape bounds
	MarginShape	[]interface{}	`json:"marginShape"`
}

// Rectangle.
type Rect  struct {

	// X coordinate
	X	float64	`json:"x"`

	// Y coordinate
	Y	float64	`json:"y"`

	// Rectangle width
	Width	float64	`json:"width"`

	// Rectangle height
	Height	float64	`json:"height"`
}
