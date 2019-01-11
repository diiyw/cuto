package tethering

const (
	
	// Request browser port binding.
	Bind = "Tethering.bind"
	
	// Request browser port unbinding.
	Unbind = "Tethering.unbind"
	
)

// Bind parameters
type BindParams struct {
	
	// Port number to bind.
	Port	int	`json:"port"`
	
}

// Bind returns
type BindReturns struct {
	
}

// Unbind parameters
type UnbindParams struct {
	
	// Port number to unbind.
	Port	int	`json:"port"`
	
}

// Unbind returns
type UnbindReturns struct {
	
}

