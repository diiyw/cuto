package layertree

// 
const LayerPaintedEvent = "LayerTree.layerPainted"
type LayerPaintedParams struct {

	// The id of the painted layer.
	LayerId 	LayerId
	// Clip rectangle.
	Clip 	interface{}}



// 
const LayerTreeDidChangeEvent = "LayerTree.layerTreeDidChange"
type LayerTreeDidChangeParams struct {

	// Layer tree, absent if not in the comspositing mode.
	Layers 	[]*Layer}

