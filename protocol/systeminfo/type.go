package systeminfo
// Describes a single graphics processor (GPU).
type GPUDevice  struct {

	// PCI ID of the GPU vendor, if available; 0 otherwise.
	VendorId	float64	`json:"vendorId"`

	// PCI ID of the GPU device, if available; 0 otherwise.
	DeviceId	float64	`json:"deviceId"`

	// Sub sys ID of the GPU, only available on Windows.
	SubSysId	float64	`json:"subSysId"`

	// Revision of the GPU, only available on Windows.
	Revision	float64	`json:"revision"`

	// String description of the GPU vendor, if the PCI ID is not available.
	VendorString	string	`json:"vendorString"`

	// String description of the GPU device, if the PCI ID is not available.
	DeviceString	string	`json:"deviceString"`

	// String description of the GPU driver vendor.
	DriverVendor	string	`json:"driverVendor"`

	// String description of the GPU driver version.
	DriverVersion	string	`json:"driverVersion"`
}

// Describes the width and height dimensions of an entity.
type Size  struct {

	// Width in pixels.
	Width	int	`json:"width"`

	// Height in pixels.
	Height	int	`json:"height"`
}

// Describes a supported video decoding profile with its associated minimum and
	// maximum resolutions.
type VideoDecodeAcceleratorCapability  struct {

	// Video codec profile that is supported, e.g. VP9 Profile 2.
	Profile	string	`json:"profile"`

	// Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution	Size	`json:"maxResolution"`

	// Minimum video dimensions in pixels supported for this |profile|.
	MinResolution	Size	`json:"minResolution"`
}

// Describes a supported video encoding profile with its associated maximum
	// resolution and maximum framerate.
type VideoEncodeAcceleratorCapability  struct {

	// Video codec profile that is supported, e.g H264 Main.
	Profile	string	`json:"profile"`

	// Maximum video dimensions in pixels supported for this |profile|.
	MaxResolution	Size	`json:"maxResolution"`

	// Maximum encoding framerate in frames per second supported for this
	// |profile|, as fraction's numerator and denominator, e.g. 24/1 fps,
	// 24000/1001 fps, etc.
	MaxFramerateNumerator	int	`json:"maxFramerateNumerator"`

	// 
	MaxFramerateDenominator	int	`json:"maxFramerateDenominator"`
}

// YUV subsampling type of the pixels of a given image.
type SubsamplingFormat string

// Image format of a given image.
type ImageType string

// Describes a supported image decoding profile with its associated minimum and
	// maximum resolutions and subsampling.
type ImageDecodeAcceleratorCapability  struct {

	// Image coded, e.g. Jpeg.
	ImageType	ImageType	`json:"imageType"`

	// Maximum supported dimensions of the image in pixels.
	MaxDimensions	Size	`json:"maxDimensions"`

	// Minimum supported dimensions of the image in pixels.
	MinDimensions	Size	`json:"minDimensions"`

	// Optional array of supported subsampling formats, e.g. 4:2:0, if known.
	Subsamplings	[]*SubsamplingFormat	`json:"subsamplings"`
}

// Provides information about the GPU(s) on the system.
type GPUInfo  struct {

	// The graphics devices on the system. Element 0 is the primary GPU.
	Devices	[]*GPUDevice	`json:"devices"`

	// An optional dictionary of additional GPU related attributes.
	AuxAttributes	interface{}	`json:"auxAttributes"`

	// An optional dictionary of graphics features and their status.
	FeatureStatus	interface{}	`json:"featureStatus"`

	// An optional array of GPU driver bug workarounds.
	DriverBugWorkarounds	[]string	`json:"driverBugWorkarounds"`

	// Supported accelerated video decoding capabilities.
	VideoDecoding	[]*VideoDecodeAcceleratorCapability	`json:"videoDecoding"`

	// Supported accelerated video encoding capabilities.
	VideoEncoding	[]*VideoEncodeAcceleratorCapability	`json:"videoEncoding"`

	// Supported accelerated image decoding capabilities.
	ImageDecoding	[]*ImageDecodeAcceleratorCapability	`json:"imageDecoding"`
}

// Represents process info.
type ProcessInfo  struct {

	// Specifies process type.
	Type	string	`json:"type"`

	// Specifies process id.
	Id	int	`json:"id"`

	// Specifies cumulative CPU usage in seconds across all threads of the
	// process since the process start.
	CpuTime	float64	`json:"cpuTime"`
}
