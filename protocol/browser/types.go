package browser


// 
type WindowID int	

// The state of the browser window.
type WindowState string	

// Browser window bounds information
type Bounds struct {
	
	// The offset from the left edge of the screen to the window in pixels.
	
	Left	int	`json:"left"`
	
	// The offset from the top edge of the screen to the window in pixels.
	
	Top	int	`json:"top"`
	
	// The window width in pixels.
	
	Width	int	`json:"width"`
	
	// The window height in pixels.
	
	Height	int	`json:"height"`
	
	// The window state. Default to normal.
	
	WindowState	WindowState	`json:"windowState"`
	
}	

// 
type PermissionType string	

// Chrome histogram bucket.
type Bucket struct {
	
	// Minimum value (inclusive).
	
	Low	int	`json:"low"`
	
	// Maximum value (exclusive).
	
	High	int	`json:"high"`
	
	// Number of samples.
	
	Count	int	`json:"count"`
	
}	

// Chrome histogram.
type Histogram struct {
	
	// Name.
	
	Name	string	`json:"name"`
	
	// Sum of sample values.
	
	Sum	int	`json:"sum"`
	
	// Total number of samples.
	
	Count	int	`json:"count"`
	
	// Buckets.
	
	Buckets	[]Bucket	`json:"buckets"`
	
}	

