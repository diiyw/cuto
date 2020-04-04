package performance
// Run-time execution metric.
type Metric  struct {

	// Metric name.
	Name	string	`json:"name"`

	// Metric value.
	Value	float64	`json:"value"`
}
