package deviceorientation

const (
	
	// Clears the overridden Device Orientation.
	ClearDeviceOrientationOverride = "DeviceOrientation.clearDeviceOrientationOverride"
	
	// Overrides the Device Orientation.
	SetDeviceOrientationOverride = "DeviceOrientation.setDeviceOrientationOverride"
	
)

// ClearDeviceOrientationOverride parameters
type ClearDeviceOrientationOverrideParams struct {
	
}

// ClearDeviceOrientationOverride returns
type ClearDeviceOrientationOverrideReturns struct {
	
}

// SetDeviceOrientationOverride parameters
type SetDeviceOrientationOverrideParams struct {
	
	// Mock alpha
	Alpha	float64	`json:"alpha"`
	
	// Mock beta
	Beta	float64	`json:"beta"`
	
	// Mock gamma
	Gamma	float64	`json:"gamma"`
	
}

// SetDeviceOrientationOverride returns
type SetDeviceOrientationOverrideReturns struct {
	
}

