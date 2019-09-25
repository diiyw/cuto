package database


// Unique identifier of Database object.
type DatabaseId string	

// Database object.
type Database struct {
	
	// Database ID.
	
	Id	DatabaseId	`json:"id"`
	
	// Database domain.
	
	Domain	string	`json:"domain"`
	
	// Database name.
	
	Name	string	`json:"name"`
	
	// Database version.
	
	Version	string	`json:"version"`
	
}	

// Database error.
type Error struct {
	
	// Error message.
	
	Message	string	`json:"message"`
	
	// Error code.
	
	Code	int	`json:"code"`
	
}	

