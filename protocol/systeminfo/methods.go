package systeminfo

const (
	
	// Returns information about the system.
	GetInfo = "SystemInfo.getInfo"
	
	// Returns information about all running processes.
	GetProcessInfo = "SystemInfo.getProcessInfo"
	
)

// GetInfo parameters
type GetInfoParams struct {
	
}

// GetInfo returns
type GetInfoReturns struct {
	
	// Information about the GPUs on the system.
	Gpu	GPUInfo	`json:"gpu"`
	
	// A platform-dependent description of the model of the machine. On Mac OS, this is, for
	// example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName	string	`json:"modelName"`
	
	// A platform-dependent description of the version of the machine. On Mac OS, this is, for
	// example, '10.1'. Will be the empty string if not supported.
	ModelVersion	string	`json:"modelVersion"`
	
	// The command line string used to launch the browser. Will be the empty string if not
	// supported.
	CommandLine	string	`json:"commandLine"`
	
}

// GetProcessInfo parameters
type GetProcessInfoParams struct {
	
}

// GetProcessInfo returns
type GetProcessInfoReturns struct {
	
	// An array of process info blocks.
	ProcessInfo	[]ProcessInfo	`json:"processInfo"`
	
}

