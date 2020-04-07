package deviceorientation

// Clears the overridden Device Orientation.
const ClearDeviceOrientationOverride = "DeviceOrientation.clearDeviceOrientationOverride"

type ClearDeviceOrientationOverrideParams struct {
}

type ClearDeviceOrientationOverrideResult struct {

}

// Overrides the Device Orientation.
const SetDeviceOrientationOverride = "DeviceOrientation.setDeviceOrientationOverride"

type SetDeviceOrientationOverrideParams struct {

	// Mock alpha
	Alpha 	float64	`json:"alpha"`

	// Mock beta
	Beta 	float64	`json:"beta"`

	// Mock gamma
	Gamma 	float64	`json:"gamma"`
}

type SetDeviceOrientationOverrideResult struct {

}