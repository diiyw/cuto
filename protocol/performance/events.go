package performance

const (
	
	// Current values of the metrics.
	MetricsEvent = "Performance.metrics"
	
)

// Current values of the metrics.
type MetricsParams struct {
	
	// Current values of the metrics.
	Metrics	[]Metric	`json:"metrics"`
	
	// Timestamp title.
	Title	string	`json:"title"`
	
}

