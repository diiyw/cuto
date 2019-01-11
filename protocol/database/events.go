package database

const (
	
	
	AddDatabaseEvent = "Database.addDatabase"
	
)


type AddDatabaseParams struct {
	
	
	Database	Database	`json:"database"`
	
}

