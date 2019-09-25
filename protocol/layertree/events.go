package layertree

import (

	"github.com/diiyw/goc/protocol/dom"

)
const (
	
	
	LayerPaintedEvent = "LayerTree.layerPainted"
	
	
	LayerTreeDidChangeEvent = "LayerTree.layerTreeDidChange"
	
)


type LayerPaintedParams struct {
	
	// The id of the painted layer.
	LayerId	LayerId	`json:"layerId"`
	
	// Clip rectangle.
	Clip	dom.Rect	`json:"clip"`
	
}


type LayerTreeDidChangeParams struct {
	
	// Layer tree, absent if not in the comspositing mode.
	Layers	[]Layer	`json:"layers"`
	
}

