package tethering

// Request browser port binding.
const Bind = "Tethering.bind"

type BindParams struct {

	// Port number to bind.
	Port 	int	`json:"port"`
}

type BindResult struct {

}

// Request browser port unbinding.
const Unbind = "Tethering.unbind"

type UnbindParams struct {

	// Port number to unbind.
	Port 	int	`json:"port"`
}

type UnbindResult struct {

}