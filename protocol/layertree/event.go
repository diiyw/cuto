package layertree

import (
	"github.com/diiyw/cuto/protocol/cdp"
)


// 
const LayerPaintedEvent = "LayerTree.layerPainted"
type LayerPaintedParams struct {

	// The id of the painted layer.
	LayerId 	LayerId
	// Clip rectangle.
	Clip 	cdp.Rect}



// 
const LayerTreeDidChangeEvent = "LayerTree.layerTreeDidChange"
type LayerTreeDidChangeParams struct {

	// Layer tree, absent if not in the comspositing mode.
	Layers 	[]*Layer}

