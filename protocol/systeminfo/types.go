package systeminfo


// Describes a single graphics processor (GPU).
type GPUDevice struct {
	
	// PCI ID of the GPU vendor, if available; 0 otherwise.
	
	VendorId	float64	`json:"vendorId"`
	
	// PCI ID of the GPU device, if available; 0 otherwise.
	
	DeviceId	float64	`json:"deviceId"`
	
	// String description of the GPU vendor, if the PCI ID is not available.
	
	VendorString	string	`json:"vendorString"`
	
	// String description of the GPU device, if the PCI ID is not available.
	
	DeviceString	string	`json:"deviceString"`
	
}	

// Provides information about the GPU(s) on the system.
type GPUInfo struct {
	
	// The graphics devices on the system. Element 0 is the primary GPU.
	
	Devices	[]GPUDevice	`json:"devices"`
	
	// An optional dictionary of additional GPU related attributes.
	
	AuxAttributes	interface{}	`json:"auxAttributes"`
	
	// An optional dictionary of graphics features and their status.
	
	FeatureStatus	interface{}	`json:"featureStatus"`
	
	// An optional array of GPU driver bug workarounds.
	
	DriverBugWorkarounds	[]string	`json:"driverBugWorkarounds"`
	
}	

// Represents process info.
type ProcessInfo struct {
	
	// Specifies process type.
	
	Type	string	`json:"type"`
	
	// Specifies process id.
	
	Id	int	`json:"id"`
	
	// Specifies cumulative CPU usage in seconds across all threads of the
	// process since the process start.
	
	CpuTime	float64	`json:"cpuTime"`
	
}	

