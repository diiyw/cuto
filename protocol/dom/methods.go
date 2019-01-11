package dom

import (

	"github.com/diiyw/goc/protocol/frame"

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Collects class names for the node with given id and all of it's child nodes.
	CollectClassNamesFromSubtree = "DOM.collectClassNamesFromSubtree"
	
	// Creates a deep copy of the specified node and places it into the target container before the
	// given anchor.
	CopyTo = "DOM.copyTo"
	
	// Describes node given its id, does not require domain to be enabled. Does not start tracking any
	// objects, can be used for automation.
	DescribeNode = "DOM.describeNode"
	
	// Disables DOM agent for the given page.
	Disable = "DOM.disable"
	
	// Discards search results from the session with the given id. `getSearchResults` should no longer
	// be called for that search.
	DiscardSearchResults = "DOM.discardSearchResults"
	
	// Enables DOM agent for the given page.
	Enable = "DOM.enable"
	
	// Focuses the given element.
	Focus = "DOM.focus"
	
	// Returns attributes for the specified node.
	GetAttributes = "DOM.getAttributes"
	
	// Returns boxes for the given node.
	GetBoxModel = "DOM.getBoxModel"
	
	// Returns quads that describe node position on the page. This method
	// might return multiple quads for inline nodes.
	GetContentQuads = "DOM.getContentQuads"
	
	// Returns the root DOM node (and optionally the subtree) to the caller.
	GetDocument = "DOM.getDocument"
	
	// Returns the root DOM node (and optionally the subtree) to the caller.
	GetFlattenedDocument = "DOM.getFlattenedDocument"
	
	// Returns node id at given location. Depending on whether DOM domain is enabled, nodeId is
	// either returned or not.
	GetNodeForLocation = "DOM.getNodeForLocation"
	
	// Returns node's HTML markup.
	GetOuterHTML = "DOM.getOuterHTML"
	
	// Returns the id of the nearest ancestor that is a relayout boundary.
	GetRelayoutBoundary = "DOM.getRelayoutBoundary"
	
	// Returns search results from given `fromIndex` to given `toIndex` from the search with the given
	// identifier.
	GetSearchResults = "DOM.getSearchResults"
	
	// Hides any highlight.
	HideHighlight = "DOM.hideHighlight"
	
	// Highlights DOM node.
	HighlightNode = "DOM.highlightNode"
	
	// Highlights given rectangle.
	HighlightRect = "DOM.highlightRect"
	
	// Marks last undoable state.
	MarkUndoableState = "DOM.markUndoableState"
	
	// Moves node into the new container, places it before the given anchor.
	MoveTo = "DOM.moveTo"
	
	// Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or
	// `cancelSearch` to end this search session.
	PerformSearch = "DOM.performSearch"
	
	// Requests that the node is sent to the caller given its path. // FIXME, use XPath
	PushNodeByPathToFrontend = "DOM.pushNodeByPathToFrontend"
	
	// Requests that a batch of nodes is sent to the caller given their backend node ids.
	PushNodesByBackendIdsToFrontend = "DOM.pushNodesByBackendIdsToFrontend"
	
	// Executes `querySelector` on a given node.
	QuerySelector = "DOM.querySelector"
	
	// Executes `querySelectorAll` on a given node.
	QuerySelectorAll = "DOM.querySelectorAll"
	
	// Re-does the last undone action.
	Redo = "DOM.redo"
	
	// Removes attribute with given name from an element with given id.
	RemoveAttribute = "DOM.removeAttribute"
	
	// Removes node with given id.
	RemoveNode = "DOM.removeNode"
	
	// Requests that children of the node with given id are returned to the caller in form of
	// `setChildNodes` events where not only immediate children are retrieved, but all children down to
	// the specified depth.
	RequestChildNodes = "DOM.requestChildNodes"
	
	// Requests that the node is sent to the caller given the JavaScript node object reference. All
	// nodes that form the path from the node to the root are also sent to the client as a series of
	// `setChildNodes` notifications.
	RequestNode = "DOM.requestNode"
	
	// Resolves the JavaScript node object for a given NodeId or BackendNodeId.
	ResolveNode = "DOM.resolveNode"
	
	// Sets attribute for an element with given id.
	SetAttributeValue = "DOM.setAttributeValue"
	
	// Sets attributes on element with given id. This method is useful when user edits some existing
	// attribute value and types in several attribute name/value pairs.
	SetAttributesAsText = "DOM.setAttributesAsText"
	
	// Sets files for the given file input element.
	SetFileInputFiles = "DOM.setFileInputFiles"
	
	// Returns file information for the given
	// File wrapper.
	GetFileInfo = "DOM.getFileInfo"
	
	// Enables console to refer to the node with given id via $x (see Command Line API for more details
	// $x functions).
	SetInspectedNode = "DOM.setInspectedNode"
	
	// Sets node name for a node with given id.
	SetNodeName = "DOM.setNodeName"
	
	// Sets node value for a node with given id.
	SetNodeValue = "DOM.setNodeValue"
	
	// Sets node HTML markup, returns new node id.
	SetOuterHTML = "DOM.setOuterHTML"
	
	// Undoes the last performed action.
	Undo = "DOM.undo"
	
	// Returns iframe node that owns iframe with the given domain.
	GetFrameOwner = "DOM.getFrameOwner"
	
)

// CollectClassNamesFromSubtree parameters
type CollectClassNamesFromSubtreeParams struct {
	
	// Id of the node to collect class names.
	NodeId	NodeId	`json:"nodeId"`
	
}

// CollectClassNamesFromSubtree returns
type CollectClassNamesFromSubtreeReturns struct {
	
	// Class name list.
	ClassNames	[]string	`json:"classNames"`
	
}

// CopyTo parameters
type CopyToParams struct {
	
	// Id of the node to copy.
	NodeId	NodeId	`json:"nodeId"`
	
	// Id of the element to drop the copy into.
	TargetNodeId	NodeId	`json:"targetNodeId"`
	
	// Drop the copy before this node (if absent, the copy becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeId	NodeId	`json:"insertBeforeNodeId"`
	
}

// CopyTo returns
type CopyToReturns struct {
	
	// Id of the node clone.
	NodeId	NodeId	`json:"nodeId"`
	
}

// DescribeNode parameters
type DescribeNodeParams struct {
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth	int	`json:"depth"`
	
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce	bool	`json:"pierce"`
	
}

// DescribeNode returns
type DescribeNodeReturns struct {
	
	// Node description.
	Node	Node	`json:"node"`
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// DiscardSearchResults parameters
type DiscardSearchResultsParams struct {
	
	// Unique search session identifier.
	SearchId	string	`json:"searchId"`
	
}

// DiscardSearchResults returns
type DiscardSearchResultsReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// Focus parameters
type FocusParams struct {
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// Focus returns
type FocusReturns struct {
	
}

// GetAttributes parameters
type GetAttributesParams struct {
	
	// Id of the node to retrieve attibutes for.
	NodeId	NodeId	`json:"nodeId"`
	
}

// GetAttributes returns
type GetAttributesReturns struct {
	
	// An interleaved array of node attribute names and values.
	Attributes	[]string	`json:"attributes"`
	
}

// GetBoxModel parameters
type GetBoxModelParams struct {
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// GetBoxModel returns
type GetBoxModelReturns struct {
	
	// Box model for the node.
	Model	BoxModel	`json:"model"`
	
}

// GetContentQuads parameters
type GetContentQuadsParams struct {
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// GetContentQuads returns
type GetContentQuadsReturns struct {
	
	// Quads that describe node layout relative to viewport.
	Quads	[]Quad	`json:"quads"`
	
}

// GetDocument parameters
type GetDocumentParams struct {
	
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth	int	`json:"depth"`
	
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce	bool	`json:"pierce"`
	
}

// GetDocument returns
type GetDocumentReturns struct {
	
	// Resulting node.
	Root	Node	`json:"root"`
	
}

// GetFlattenedDocument parameters
type GetFlattenedDocumentParams struct {
	
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth	int	`json:"depth"`
	
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce	bool	`json:"pierce"`
	
}

// GetFlattenedDocument returns
type GetFlattenedDocumentReturns struct {
	
	// Resulting node.
	Nodes	[]Node	`json:"nodes"`
	
}

// GetNodeForLocation parameters
type GetNodeForLocationParams struct {
	
	// X coordinate.
	X	int	`json:"x"`
	
	// Y coordinate.
	Y	int	`json:"y"`
	
	// False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM	bool	`json:"includeUserAgentShadowDOM"`
	
}

// GetNodeForLocation returns
type GetNodeForLocationReturns struct {
	
	// Resulting node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// Id of the node at given coordinates, only when enabled.
	NodeId	NodeId	`json:"nodeId"`
	
}

// GetOuterHTML parameters
type GetOuterHTMLParams struct {
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// GetOuterHTML returns
type GetOuterHTMLReturns struct {
	
	// Outer HTML markup.
	OuterHTML	string	`json:"outerHTML"`
	
}

// GetRelayoutBoundary parameters
type GetRelayoutBoundaryParams struct {
	
	// Id of the node.
	NodeId	NodeId	`json:"nodeId"`
	
}

// GetRelayoutBoundary returns
type GetRelayoutBoundaryReturns struct {
	
	// Relayout boundary node id for the given node.
	NodeId	NodeId	`json:"nodeId"`
	
}

// GetSearchResults parameters
type GetSearchResultsParams struct {
	
	// Unique search session identifier.
	SearchId	string	`json:"searchId"`
	
	// Start index of the search result to be returned.
	FromIndex	int	`json:"fromIndex"`
	
	// End index of the search result to be returned.
	ToIndex	int	`json:"toIndex"`
	
}

// GetSearchResults returns
type GetSearchResultsReturns struct {
	
	// Ids of the search result nodes.
	NodeIds	[]NodeId	`json:"nodeIds"`
	
}

// HideHighlight parameters
type HideHighlightParams struct {
	
}

// HideHighlight returns
type HideHighlightReturns struct {
	
}

// HighlightNode parameters
type HighlightNodeParams struct {
	
}

// HighlightNode returns
type HighlightNodeReturns struct {
	
}

// HighlightRect parameters
type HighlightRectParams struct {
	
}

// HighlightRect returns
type HighlightRectReturns struct {
	
}

// MarkUndoableState parameters
type MarkUndoableStateParams struct {
	
}

// MarkUndoableState returns
type MarkUndoableStateReturns struct {
	
}

// MoveTo parameters
type MoveToParams struct {
	
	// Id of the node to move.
	NodeId	NodeId	`json:"nodeId"`
	
	// Id of the element to drop the moved node into.
	TargetNodeId	NodeId	`json:"targetNodeId"`
	
	// Drop node before this one (if absent, the moved node becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeId	NodeId	`json:"insertBeforeNodeId"`
	
}

// MoveTo returns
type MoveToReturns struct {
	
	// New id of the moved node.
	NodeId	NodeId	`json:"nodeId"`
	
}

// PerformSearch parameters
type PerformSearchParams struct {
	
	// Plain text or query selector or XPath search query.
	Query	string	`json:"query"`
	
	// True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM	bool	`json:"includeUserAgentShadowDOM"`
	
}

// PerformSearch returns
type PerformSearchReturns struct {
	
	// Unique search session identifier.
	SearchId	string	`json:"searchId"`
	
	// Number of search results.
	ResultCount	int	`json:"resultCount"`
	
}

// PushNodeByPathToFrontend parameters
type PushNodeByPathToFrontendParams struct {
	
	// Path to node in the proprietary format.
	Path	string	`json:"path"`
	
}

// PushNodeByPathToFrontend returns
type PushNodeByPathToFrontendReturns struct {
	
	// Id of the node for given path.
	NodeId	NodeId	`json:"nodeId"`
	
}

// PushNodesByBackendIdsToFrontend parameters
type PushNodesByBackendIdsToFrontendParams struct {
	
	// The array of backend node ids.
	BackendNodeIds	[]BackendNodeId	`json:"backendNodeIds"`
	
}

// PushNodesByBackendIdsToFrontend returns
type PushNodesByBackendIdsToFrontendReturns struct {
	
	// The array of ids of pushed nodes that correspond to the backend ids specified in
	// backendNodeIds.
	NodeIds	[]NodeId	`json:"nodeIds"`
	
}

// QuerySelector parameters
type QuerySelectorParams struct {
	
	// Id of the node to query upon.
	NodeId	NodeId	`json:"nodeId"`
	
	// Selector string.
	Selector	string	`json:"selector"`
	
}

// QuerySelector returns
type QuerySelectorReturns struct {
	
	// Query selector result.
	NodeId	NodeId	`json:"nodeId"`
	
}

// QuerySelectorAll parameters
type QuerySelectorAllParams struct {
	
	// Id of the node to query upon.
	NodeId	NodeId	`json:"nodeId"`
	
	// Selector string.
	Selector	string	`json:"selector"`
	
}

// QuerySelectorAll returns
type QuerySelectorAllReturns struct {
	
	// Query selector result.
	NodeIds	[]NodeId	`json:"nodeIds"`
	
}

// Redo parameters
type RedoParams struct {
	
}

// Redo returns
type RedoReturns struct {
	
}

// RemoveAttribute parameters
type RemoveAttributeParams struct {
	
	// Id of the element to remove attribute from.
	NodeId	NodeId	`json:"nodeId"`
	
	// Name of the attribute to remove.
	Name	string	`json:"name"`
	
}

// RemoveAttribute returns
type RemoveAttributeReturns struct {
	
}

// RemoveNode parameters
type RemoveNodeParams struct {
	
	// Id of the node to remove.
	NodeId	NodeId	`json:"nodeId"`
	
}

// RemoveNode returns
type RemoveNodeReturns struct {
	
}

// RequestChildNodes parameters
type RequestChildNodesParams struct {
	
	// Id of the node to get children for.
	NodeId	NodeId	`json:"nodeId"`
	
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth	int	`json:"depth"`
	
	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree
	// (default is false).
	Pierce	bool	`json:"pierce"`
	
}

// RequestChildNodes returns
type RequestChildNodesReturns struct {
	
}

// RequestNode parameters
type RequestNodeParams struct {
	
	// JavaScript object id to convert into node.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// RequestNode returns
type RequestNodeReturns struct {
	
	// Node id for given object.
	NodeId	NodeId	`json:"nodeId"`
	
}

// ResolveNode parameters
type ResolveNodeParams struct {
	
	// Id of the node to resolve.
	NodeId	NodeId	`json:"nodeId"`
	
	// Backend identifier of the node to resolve.
	BackendNodeId		`json:"backendNodeId"`
	
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup	string	`json:"objectGroup"`
	
}

// ResolveNode returns
type ResolveNodeReturns struct {
	
	// JavaScript object wrapper for given node.
	Object	runtime.RemoteObject	`json:"object"`
	
}

// SetAttributeValue parameters
type SetAttributeValueParams struct {
	
	// Id of the element to set attribute for.
	NodeId	NodeId	`json:"nodeId"`
	
	// Attribute name.
	Name	string	`json:"name"`
	
	// Attribute value.
	Value	string	`json:"value"`
	
}

// SetAttributeValue returns
type SetAttributeValueReturns struct {
	
}

// SetAttributesAsText parameters
type SetAttributesAsTextParams struct {
	
	// Id of the element to set attributes for.
	NodeId	NodeId	`json:"nodeId"`
	
	// Text with a number of attributes. Will parse this text using HTML parser.
	Text	string	`json:"text"`
	
	// Attribute name to replace with new attributes derived from text in case text parsed
	// successfully.
	Name	string	`json:"name"`
	
}

// SetAttributesAsText returns
type SetAttributesAsTextReturns struct {
	
}

// SetFileInputFiles parameters
type SetFileInputFilesParams struct {
	
	// Array of file paths to set.
	Files	[]string	`json:"files"`
	
	// Identifier of the node.
	NodeId	NodeId	`json:"nodeId"`
	
	// Identifier of the backend node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// SetFileInputFiles returns
type SetFileInputFilesReturns struct {
	
}

// GetFileInfo parameters
type GetFileInfoParams struct {
	
	// JavaScript object id of the node wrapper.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// GetFileInfo returns
type GetFileInfoReturns struct {
	
	
	Path	string	`json:"path"`
	
}

// SetInspectedNode parameters
type SetInspectedNodeParams struct {
	
	// DOM node id to be accessible by means of $x command line API.
	NodeId	NodeId	`json:"nodeId"`
	
}

// SetInspectedNode returns
type SetInspectedNodeReturns struct {
	
}

// SetNodeName parameters
type SetNodeNameParams struct {
	
	// Id of the node to set name for.
	NodeId	NodeId	`json:"nodeId"`
	
	// New node's name.
	Name	string	`json:"name"`
	
}

// SetNodeName returns
type SetNodeNameReturns struct {
	
	// New node's id.
	NodeId	NodeId	`json:"nodeId"`
	
}

// SetNodeValue parameters
type SetNodeValueParams struct {
	
	// Id of the node to set value for.
	NodeId	NodeId	`json:"nodeId"`
	
	// New node's value.
	Value	string	`json:"value"`
	
}

// SetNodeValue returns
type SetNodeValueReturns struct {
	
}

// SetOuterHTML parameters
type SetOuterHTMLParams struct {
	
	// Id of the node to set markup for.
	NodeId	NodeId	`json:"nodeId"`
	
	// Outer HTML markup to set.
	OuterHTML	string	`json:"outerHTML"`
	
}

// SetOuterHTML returns
type SetOuterHTMLReturns struct {
	
}

// Undo parameters
type UndoParams struct {
	
}

// Undo returns
type UndoReturns struct {
	
}

// GetFrameOwner parameters
type GetFrameOwnerParams struct {
	
	
	FrameId	frame.FrameId	`json:"frameId"`
	
}

// GetFrameOwner returns
type GetFrameOwnerReturns struct {
	
	// Resulting node.
	BackendNodeId	BackendNodeId	`json:"backendNodeId"`
	
	// Id of the node at given coordinates, only when enabled.
	NodeId	NodeId	`json:"nodeId"`
	
}

