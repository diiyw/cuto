package headlessexperimental


// Encoding options for a screenshot.
type ScreenshotParams struct {
	
	// Image compression format (defaults to png).
	// Possible value: jpeg,png,
	Format	string	`json:"format"`
	
	// Compression quality from range [0..100] (jpeg only).
	
	Quality	int	`json:"quality"`
	
}	

