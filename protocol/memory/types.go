package memory


// Memory pressure level.
type PressureLevel string	

// Heap profile sample.
type SamplingProfileNode struct {
	
	// Size of the sampled allocation.
	
	Size	float64	`json:"size"`
	
	// Total bytes attributed to this sample.
	
	Total	float64	`json:"total"`
	
	// Execution stack at the point of allocation.
	
	Stack	[]string	`json:"stack"`
	
}	

// Array of heap profile samples.
type SamplingProfile struct {
	
	
	
	Samples	[]SamplingProfileNode	`json:"samples"`
	
	
	
	Modules	[]Module	`json:"modules"`
	
}	

// Executable module information
type Module struct {
	
	// Name of the module.
	
	Name	string	`json:"name"`
	
	// UUID of the module.
	
	Uuid	string	`json:"uuid"`
	
	// Base address where the module is loaded into memory. Encoded as a decimal
	// or hexadecimal (0x prefixed) string.
	
	BaseAddress	string	`json:"baseAddress"`
	
	// Size of the module in bytes.
	
	Size	float64	`json:"size"`
	
}	

