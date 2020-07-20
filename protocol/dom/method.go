package dom

import (
	"github.com/diiyw/cuto/protocol/runtime"
	"github.com/diiyw/cuto/protocol/cdp"
)


// Collects class names for the node with given id and all of it's child nodes.
const CollectClassNamesFromSubtree = "DOM.collectClassNamesFromSubtree"

type CollectClassNamesFromSubtreeParams struct {

	// Id of the node to collect class names.
	NodeId 	NodeId	`json:"nodeId"`
}

type CollectClassNamesFromSubtreeResult struct {

	// Class name list.
	ClassNames 	[]string	`json:"classNames"`
}

// Creates a deep copy of the specified node and places it into the target container before the
// given anchor.
const CopyTo = "DOM.copyTo"

type CopyToParams struct {

	// Id of the node to copy.
	NodeId 	NodeId	`json:"nodeId"`

	// Id of the element to drop the copy into.
	TargetNodeId 	NodeId	`json:"targetNodeId"`

	// Drop the copy before this node (if absent, the copy becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeId 	NodeId	`json:"insertBeforeNodeId,omitempty"`
}

type CopyToResult struct {

	// Id of the node clone.
	NodeId 	NodeId	`json:"nodeId"`
}

// Describes node given its id, does not require domain to be enabled. Does not start tracking any
// objects, can be used for automation.
const DescribeNode = "DOM.describeNode"

type DescribeNodeParams struct {

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`

	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth 	int	`json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce 	bool	`json:"pierce,omitempty"`
}

type DescribeNodeResult struct {

	// Node description.
	Node 	Node	`json:"node"`
}

// Disables DOM agent for the given page.
const Disable = "DOM.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Discards search results from the session with the given id. `getSearchResults` should no longer
// be called for that search.
const DiscardSearchResults = "DOM.discardSearchResults"

type DiscardSearchResultsParams struct {

	// Unique search session identifier.
	SearchId 	string	`json:"searchId"`
}

type DiscardSearchResultsResult struct {

}

// Enables DOM agent for the given page.
const Enable = "DOM.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Focuses the given element.
const Focus = "DOM.focus"

type FocusParams struct {

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`
}

type FocusResult struct {

}

// Returns attributes for the specified node.
const GetAttributes = "DOM.getAttributes"

type GetAttributesParams struct {

	// Id of the node to retrieve attibutes for.
	NodeId 	NodeId	`json:"nodeId"`
}

type GetAttributesResult struct {

	// An interleaved array of node attribute names and values.
	Attributes 	[]string	`json:"attributes"`
}

// Returns boxes for the given node.
const GetBoxModel = "DOM.getBoxModel"

type GetBoxModelParams struct {

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`
}

type GetBoxModelResult struct {

	// Box model for the node.
	Model 	BoxModel	`json:"model"`
}

// Returns quads that describe node position on the page. This method
// might return multiple quads for inline nodes.
const GetContentQuads = "DOM.getContentQuads"

type GetContentQuadsParams struct {

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`
}

type GetContentQuadsResult struct {

	// Quads that describe node layout relative to viewport.
	Quads 	[]*Quad	`json:"quads"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
const GetDocument = "DOM.getDocument"

type GetDocumentParams struct {

	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth 	int	`json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce 	bool	`json:"pierce,omitempty"`
}

type GetDocumentResult struct {

	// Resulting node.
	Root 	Node	`json:"root"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
const GetFlattenedDocument = "DOM.getFlattenedDocument"

type GetFlattenedDocumentParams struct {

	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth 	int	`json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false).
	Pierce 	bool	`json:"pierce,omitempty"`
}

type GetFlattenedDocumentResult struct {

	// Resulting node.
	Nodes 	[]*Node	`json:"nodes"`
}

// Returns node id at given location. Depending on whether DOM domain is enabled, nodeId is
// either returned or not.
const GetNodeForLocation = "DOM.getNodeForLocation"

type GetNodeForLocationParams struct {

	// X coordinate.
	X 	int	`json:"x"`

	// Y coordinate.
	Y 	int	`json:"y"`

	// False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM 	bool	`json:"includeUserAgentShadowDOM,omitempty"`

	// Whether to ignore pointer-events: none on elements and hit test them.
	IgnorePointerEventsNone 	bool	`json:"ignorePointerEventsNone,omitempty"`
}

type GetNodeForLocationResult struct {

	// Resulting node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId"`
	// Frame this node belongs to.
	FrameId 	cdp.FrameId	`json:"frameId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeId 	NodeId	`json:"nodeId"`
}

// Returns node's HTML markup.
const GetOuterHTML = "DOM.getOuterHTML"

type GetOuterHTMLParams struct {

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`
}

type GetOuterHTMLResult struct {

	// Outer HTML markup.
	OuterHTML 	string	`json:"outerHTML"`
}

// Returns the id of the nearest ancestor that is a relayout boundary.
const GetRelayoutBoundary = "DOM.getRelayoutBoundary"

type GetRelayoutBoundaryParams struct {

	// Id of the node.
	NodeId 	NodeId	`json:"nodeId"`
}

type GetRelayoutBoundaryResult struct {

	// Relayout boundary node id for the given node.
	NodeId 	NodeId	`json:"nodeId"`
}

// Returns search results from given `fromIndex` to given `toIndex` from the search with the given
// identifier.
const GetSearchResults = "DOM.getSearchResults"

type GetSearchResultsParams struct {

	// Unique search session identifier.
	SearchId 	string	`json:"searchId"`

	// Start index of the search result to be returned.
	FromIndex 	int	`json:"fromIndex"`

	// End index of the search result to be returned.
	ToIndex 	int	`json:"toIndex"`
}

type GetSearchResultsResult struct {

	// Ids of the search result nodes.
	NodeIds 	[]*NodeId	`json:"nodeIds"`
}

// Hides any highlight.
const HideHighlight = "DOM.hideHighlight"

type HideHighlightParams struct {
}

type HideHighlightResult struct {

}

// Highlights DOM node.
const HighlightNode = "DOM.highlightNode"

type HighlightNodeParams struct {
}

type HighlightNodeResult struct {

}

// Highlights given rectangle.
const HighlightRect = "DOM.highlightRect"

type HighlightRectParams struct {
}

type HighlightRectResult struct {

}

// Marks last undoable state.
const MarkUndoableState = "DOM.markUndoableState"

type MarkUndoableStateParams struct {
}

type MarkUndoableStateResult struct {

}

// Moves node into the new container, places it before the given anchor.
const MoveTo = "DOM.moveTo"

type MoveToParams struct {

	// Id of the node to move.
	NodeId 	NodeId	`json:"nodeId"`

	// Id of the element to drop the moved node into.
	TargetNodeId 	NodeId	`json:"targetNodeId"`

	// Drop node before this one (if absent, the moved node becomes the last child of
	// `targetNodeId`).
	InsertBeforeNodeId 	NodeId	`json:"insertBeforeNodeId,omitempty"`
}

type MoveToResult struct {

	// New id of the moved node.
	NodeId 	NodeId	`json:"nodeId"`
}

// Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or
// `cancelSearch` to end this search session.
const PerformSearch = "DOM.performSearch"

type PerformSearchParams struct {

	// Plain text or query selector or XPath search query.
	Query 	string	`json:"query"`

	// True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM 	bool	`json:"includeUserAgentShadowDOM,omitempty"`
}

type PerformSearchResult struct {

	// Unique search session identifier.
	SearchId 	string	`json:"searchId"`
	// Number of search results.
	ResultCount 	int	`json:"resultCount"`
}

// Requests that the node is sent to the caller given its path. // FIXME, use XPath
const PushNodeByPathToFrontend = "DOM.pushNodeByPathToFrontend"

type PushNodeByPathToFrontendParams struct {

	// Path to node in the proprietary format.
	Path 	string	`json:"path"`
}

type PushNodeByPathToFrontendResult struct {

	// Id of the node for given path.
	NodeId 	NodeId	`json:"nodeId"`
}

// Requests that a batch of nodes is sent to the caller given their backend node ids.
const PushNodesByBackendIdsToFrontend = "DOM.pushNodesByBackendIdsToFrontend"

type PushNodesByBackendIdsToFrontendParams struct {

	// The array of backend node ids.
	BackendNodeIds 	[]*BackendNodeId	`json:"backendNodeIds"`
}

type PushNodesByBackendIdsToFrontendResult struct {

	// The array of ids of pushed nodes that correspond to the backend ids specified in
	// backendNodeIds.
	NodeIds 	[]*NodeId	`json:"nodeIds"`
}

// Executes `querySelector` on a given node.
const QuerySelector = "DOM.querySelector"

type QuerySelectorParams struct {

	// Id of the node to query upon.
	NodeId 	NodeId	`json:"nodeId"`

	// Selector string.
	Selector 	string	`json:"selector"`
}

type QuerySelectorResult struct {

	// Query selector result.
	NodeId 	NodeId	`json:"nodeId"`
}

// Executes `querySelectorAll` on a given node.
const QuerySelectorAll = "DOM.querySelectorAll"

type QuerySelectorAllParams struct {

	// Id of the node to query upon.
	NodeId 	NodeId	`json:"nodeId"`

	// Selector string.
	Selector 	string	`json:"selector"`
}

type QuerySelectorAllResult struct {

	// Query selector result.
	NodeIds 	[]*NodeId	`json:"nodeIds"`
}

// Re-does the last undone action.
const Redo = "DOM.redo"

type RedoParams struct {
}

type RedoResult struct {

}

// Removes attribute with given name from an element with given id.
const RemoveAttribute = "DOM.removeAttribute"

type RemoveAttributeParams struct {

	// Id of the element to remove attribute from.
	NodeId 	NodeId	`json:"nodeId"`

	// Name of the attribute to remove.
	Name 	string	`json:"name"`
}

type RemoveAttributeResult struct {

}

// Removes node with given id.
const RemoveNode = "DOM.removeNode"

type RemoveNodeParams struct {

	// Id of the node to remove.
	NodeId 	NodeId	`json:"nodeId"`
}

type RemoveNodeResult struct {

}

// Requests that children of the node with given id are returned to the caller in form of
// `setChildNodes` events where not only immediate children are retrieved, but all children down to
// the specified depth.
const RequestChildNodes = "DOM.requestChildNodes"

type RequestChildNodesParams struct {

	// Id of the node to get children for.
	NodeId 	NodeId	`json:"nodeId"`

	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth 	int	`json:"depth,omitempty"`

	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree
	// (default is false).
	Pierce 	bool	`json:"pierce,omitempty"`
}

type RequestChildNodesResult struct {

}

// Requests that the node is sent to the caller given the JavaScript node object reference. All
// nodes that form the path from the node to the root are also sent to the client as a series of
// `setChildNodes` notifications.
const RequestNode = "DOM.requestNode"

type RequestNodeParams struct {

	// JavaScript object id to convert into node.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId"`
}

type RequestNodeResult struct {

	// Node id for given object.
	NodeId 	NodeId	`json:"nodeId"`
}

// Resolves the JavaScript node object for a given NodeId or BackendNodeId.
const ResolveNode = "DOM.resolveNode"

type ResolveNodeParams struct {

	// Id of the node to resolve.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Backend identifier of the node to resolve.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`

	// Execution context in which to resolve the node.
	ExecutionContextId 	runtime.ExecutionContextId	`json:"executionContextId,omitempty"`
}

type ResolveNodeResult struct {

	// JavaScript object wrapper for given node.
	Object 	runtime.RemoteObject	`json:"object"`
}

// Sets attribute for an element with given id.
const SetAttributeValue = "DOM.setAttributeValue"

type SetAttributeValueParams struct {

	// Id of the element to set attribute for.
	NodeId 	NodeId	`json:"nodeId"`

	// Attribute name.
	Name 	string	`json:"name"`

	// Attribute value.
	Value 	string	`json:"value"`
}

type SetAttributeValueResult struct {

}

// Sets attributes on element with given id. This method is useful when user edits some existing
// attribute value and types in several attribute name/value pairs.
const SetAttributesAsText = "DOM.setAttributesAsText"

type SetAttributesAsTextParams struct {

	// Id of the element to set attributes for.
	NodeId 	NodeId	`json:"nodeId"`

	// Text with a number of attributes. Will parse this text using HTML parser.
	Text 	string	`json:"text"`

	// Attribute name to replace with new attributes derived from text in case text parsed
	// successfully.
	Name 	string	`json:"name,omitempty"`
}

type SetAttributesAsTextResult struct {

}

// Sets files for the given file input element.
const SetFileInputFiles = "DOM.setFileInputFiles"

type SetFileInputFilesParams struct {

	// Array of file paths to set.
	Files 	[]string	`json:"files"`

	// Identifier of the node.
	NodeId 	NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`
}

type SetFileInputFilesResult struct {

}

// Sets if stack traces should be captured for Nodes. See `Node.getNodeStackTraces`. Default is disabled.
const SetNodeStackTracesEnabled = "DOM.setNodeStackTracesEnabled"

type SetNodeStackTracesEnabledParams struct {

	// Enable or disable.
	Enable 	bool	`json:"enable"`
}

type SetNodeStackTracesEnabledResult struct {

}

// Gets stack traces associated with a Node. As of now, only provides stack trace for Node creation.
const GetNodeStackTraces = "DOM.getNodeStackTraces"

type GetNodeStackTracesParams struct {

	// Id of the node to get stack traces for.
	NodeId 	NodeId	`json:"nodeId"`
}

type GetNodeStackTracesResult struct {

	// Creation stack trace, if available.
	Creation 	runtime.StackTrace	`json:"creation"`
}

// Returns file information for the given
// File wrapper.
const GetFileInfo = "DOM.getFileInfo"

type GetFileInfoParams struct {

	// JavaScript object id of the node wrapper.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId"`
}

type GetFileInfoResult struct {

	// 
	Path 	string	`json:"path"`
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details
// $x functions).
const SetInspectedNode = "DOM.setInspectedNode"

type SetInspectedNodeParams struct {

	// DOM node id to be accessible by means of $x command line API.
	NodeId 	NodeId	`json:"nodeId"`
}

type SetInspectedNodeResult struct {

}

// Sets node name for a node with given id.
const SetNodeName = "DOM.setNodeName"

type SetNodeNameParams struct {

	// Id of the node to set name for.
	NodeId 	NodeId	`json:"nodeId"`

	// New node's name.
	Name 	string	`json:"name"`
}

type SetNodeNameResult struct {

	// New node's id.
	NodeId 	NodeId	`json:"nodeId"`
}

// Sets node value for a node with given id.
const SetNodeValue = "DOM.setNodeValue"

type SetNodeValueParams struct {

	// Id of the node to set value for.
	NodeId 	NodeId	`json:"nodeId"`

	// New node's value.
	Value 	string	`json:"value"`
}

type SetNodeValueResult struct {

}

// Sets node HTML markup, returns new node id.
const SetOuterHTML = "DOM.setOuterHTML"

type SetOuterHTMLParams struct {

	// Id of the node to set markup for.
	NodeId 	NodeId	`json:"nodeId"`

	// Outer HTML markup to set.
	OuterHTML 	string	`json:"outerHTML"`
}

type SetOuterHTMLResult struct {

}

// Undoes the last performed action.
const Undo = "DOM.undo"

type UndoParams struct {
}

type UndoResult struct {

}

// Returns iframe node that owns iframe with the given domain.
const GetFrameOwner = "DOM.getFrameOwner"

type GetFrameOwnerParams struct {

	// 
	FrameId 	cdp.FrameId	`json:"frameId"`
}

type GetFrameOwnerResult struct {

	// Resulting node.
	BackendNodeId 	BackendNodeId	`json:"backendNodeId"`
	// Id of the node at given coordinates, only when enabled and requested document.
	NodeId 	NodeId	`json:"nodeId"`
}