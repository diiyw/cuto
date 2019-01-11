package dom

const (
	
	// Fired when `Element`'s attribute is modified.
	AttributeModifiedEvent = "DOM.attributeModified"
	
	// Fired when `Element`'s attribute is removed.
	AttributeRemovedEvent = "DOM.attributeRemoved"
	
	// Mirrors `DOMCharacterDataModified` event.
	CharacterDataModifiedEvent = "DOM.characterDataModified"
	
	// Fired when `Container`'s child node count has changed.
	ChildNodeCountUpdatedEvent = "DOM.childNodeCountUpdated"
	
	// Mirrors `DOMNodeInserted` event.
	ChildNodeInsertedEvent = "DOM.childNodeInserted"
	
	// Mirrors `DOMNodeRemoved` event.
	ChildNodeRemovedEvent = "DOM.childNodeRemoved"
	
	// Called when distrubution is changed.
	DistributedNodesUpdatedEvent = "DOM.distributedNodesUpdated"
	
	// Fired when `Document` has been totally updated. Node ids are no longer valid.
	DocumentUpdatedEvent = "DOM.documentUpdated"
	
	// Fired when `Element`'s inline style is modified via a CSS property modification.
	InlineStyleInvalidatedEvent = "DOM.inlineStyleInvalidated"
	
	// Called when a pseudo element is added to an element.
	PseudoElementAddedEvent = "DOM.pseudoElementAdded"
	
	// Called when a pseudo element is removed from an element.
	PseudoElementRemovedEvent = "DOM.pseudoElementRemoved"
	
	// Fired when backend wants to provide client with the missing DOM structure. This happens upon
	// most of the calls requesting node ids.
	SetChildNodesEvent = "DOM.setChildNodes"
	
	// Called when shadow root is popped from the element.
	ShadowRootPoppedEvent = "DOM.shadowRootPopped"
	
	// Called when shadow root is pushed into the element.
	ShadowRootPushedEvent = "DOM.shadowRootPushed"
	
)

// Fired when `Element`'s attribute is modified.
type AttributeModifiedParams struct {
	
	// Id of the node that has changed.
	NodeId	NodeId	`json:"nodeId"`
	
	// Attribute name.
	Name	string	`json:"name"`
	
	// Attribute value.
	Value	string	`json:"value"`
	
}

// Fired when `Element`'s attribute is removed.
type AttributeRemovedParams struct {
	
	// Id of the node that has changed.
	NodeId	NodeId	`json:"nodeId"`
	
	// A ttribute name.
	Name	string	`json:"name"`
	
}

// Mirrors `DOMCharacterDataModified` event.
type CharacterDataModifiedParams struct {
	
	// Id of the node that has changed.
	NodeId	NodeId	`json:"nodeId"`
	
	// New text value.
	CharacterData	string	`json:"characterData"`
	
}

// Fired when `Container`'s child node count has changed.
type ChildNodeCountUpdatedParams struct {
	
	// Id of the node that has changed.
	NodeId	NodeId	`json:"nodeId"`
	
	// New node count.
	ChildNodeCount	int	`json:"childNodeCount"`
	
}

// Mirrors `DOMNodeInserted` event.
type ChildNodeInsertedParams struct {
	
	// Id of the node that has changed.
	ParentNodeId	NodeId	`json:"parentNodeId"`
	
	// If of the previous siblint.
	PreviousNodeId	NodeId	`json:"previousNodeId"`
	
	// Inserted node data.
	Node	Node	`json:"node"`
	
}

// Mirrors `DOMNodeRemoved` event.
type ChildNodeRemovedParams struct {
	
	// Parent id.
	ParentNodeId	NodeId	`json:"parentNodeId"`
	
	// Id of the node that has been removed.
	NodeId	NodeId	`json:"nodeId"`
	
}

// Called when distrubution is changed.
type DistributedNodesUpdatedParams struct {
	
	// Insertion point where distrubuted nodes were updated.
	InsertionPointId	NodeId	`json:"insertionPointId"`
	
	// Distributed nodes for given insertion point.
	DistributedNodes	[]BackendNode	`json:"distributedNodes"`
	
}

// Fired when `Document` has been totally updated. Node ids are no longer valid.
type DocumentUpdatedParams struct {
	
}

// Fired when `Element`'s inline style is modified via a CSS property modification.
type InlineStyleInvalidatedParams struct {
	
	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds	[]NodeId	`json:"nodeIds"`
	
}

// Called when a pseudo element is added to an element.
type PseudoElementAddedParams struct {
	
	// Pseudo element's parent element id.
	ParentId	NodeId	`json:"parentId"`
	
	// The added pseudo element.
	PseudoElement	Node	`json:"pseudoElement"`
	
}

// Called when a pseudo element is removed from an element.
type PseudoElementRemovedParams struct {
	
	// Pseudo element's parent element id.
	ParentId	NodeId	`json:"parentId"`
	
	// The removed pseudo element id.
	PseudoElementId	NodeId	`json:"pseudoElementId"`
	
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon
	// most of the calls requesting node ids.
type SetChildNodesParams struct {
	
	// Parent node id to populate with children.
	ParentId	NodeId	`json:"parentId"`
	
	// Child nodes array.
	Nodes	[]Node	`json:"nodes"`
	
}

// Called when shadow root is popped from the element.
type ShadowRootPoppedParams struct {
	
	// Host element id.
	HostId	NodeId	`json:"hostId"`
	
	// Shadow root id.
	RootId	NodeId	`json:"rootId"`
	
}

// Called when shadow root is pushed into the element.
type ShadowRootPushedParams struct {
	
	// Host element id.
	HostId	NodeId	`json:"hostId"`
	
	// Shadow root.
	Root	Node	`json:"root"`
	
}

