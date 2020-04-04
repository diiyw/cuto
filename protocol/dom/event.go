package dom

// Fired when `Element`'s attribute is modified.
const AttributeModifiedEvent = "DOM.attributeModified"
type AttributeModifiedParams struct {

	// Id of the node that has changed.
	NodeId 	NodeId
	// Attribute name.
	Name 	string
	// Attribute value.
	Value 	string}



// Fired when `Element`'s attribute is removed.
const AttributeRemovedEvent = "DOM.attributeRemoved"
type AttributeRemovedParams struct {

	// Id of the node that has changed.
	NodeId 	NodeId
	// A ttribute name.
	Name 	string}



// Mirrors `DOMCharacterDataModified` event.
const CharacterDataModifiedEvent = "DOM.characterDataModified"
type CharacterDataModifiedParams struct {

	// Id of the node that has changed.
	NodeId 	NodeId
	// New text value.
	CharacterData 	string}



// Fired when `Container`'s child node count has changed.
const ChildNodeCountUpdatedEvent = "DOM.childNodeCountUpdated"
type ChildNodeCountUpdatedParams struct {

	// Id of the node that has changed.
	NodeId 	NodeId
	// New node count.
	ChildNodeCount 	int}



// Mirrors `DOMNodeInserted` event.
const ChildNodeInsertedEvent = "DOM.childNodeInserted"
type ChildNodeInsertedParams struct {

	// Id of the node that has changed.
	ParentNodeId 	NodeId
	// If of the previous siblint.
	PreviousNodeId 	NodeId
	// Inserted node data.
	Node 	Node}



// Mirrors `DOMNodeRemoved` event.
const ChildNodeRemovedEvent = "DOM.childNodeRemoved"
type ChildNodeRemovedParams struct {

	// Parent id.
	ParentNodeId 	NodeId
	// Id of the node that has been removed.
	NodeId 	NodeId}



// Called when distrubution is changed.
const DistributedNodesUpdatedEvent = "DOM.distributedNodesUpdated"
type DistributedNodesUpdatedParams struct {

	// Insertion point where distrubuted nodes were updated.
	InsertionPointId 	NodeId
	// Distributed nodes for given insertion point.
	DistributedNodes 	[]*BackendNode}



// Fired when `Document` has been totally updated. Node ids are no longer valid.
const DocumentUpdatedEvent = "DOM.documentUpdated"
type DocumentUpdatedParams struct {
}



// Fired when `Element`'s inline style is modified via a CSS property modification.
const InlineStyleInvalidatedEvent = "DOM.inlineStyleInvalidated"
type InlineStyleInvalidatedParams struct {

	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds 	[]*NodeId}



// Called when a pseudo element is added to an element.
const PseudoElementAddedEvent = "DOM.pseudoElementAdded"
type PseudoElementAddedParams struct {

	// Pseudo element's parent element id.
	ParentId 	NodeId
	// The added pseudo element.
	PseudoElement 	Node}



// Called when a pseudo element is removed from an element.
const PseudoElementRemovedEvent = "DOM.pseudoElementRemoved"
type PseudoElementRemovedParams struct {

	// Pseudo element's parent element id.
	ParentId 	NodeId
	// The removed pseudo element id.
	PseudoElementId 	NodeId}



// Fired when backend wants to provide client with the missing DOM structure. This happens upon
// most of the calls requesting node ids.
const SetChildNodesEvent = "DOM.setChildNodes"
type SetChildNodesParams struct {

	// Parent node id to populate with children.
	ParentId 	NodeId
	// Child nodes array.
	Nodes 	[]*Node}



// Called when shadow root is popped from the element.
const ShadowRootPoppedEvent = "DOM.shadowRootPopped"
type ShadowRootPoppedParams struct {

	// Host element id.
	HostId 	NodeId
	// Shadow root id.
	RootId 	NodeId}



// Called when shadow root is pushed into the element.
const ShadowRootPushedEvent = "DOM.shadowRootPushed"
type ShadowRootPushedParams struct {

	// Host element id.
	HostId 	NodeId
	// Shadow root.
	Root 	Node}

