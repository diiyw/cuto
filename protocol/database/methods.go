package database

const (
	
	// Disables database tracking, prevents database events from being sent to the client.
	Disable = "Database.disable"
	
	// Enables database tracking, database events will now be delivered to the client.
	Enable = "Database.enable"
	
	
	ExecuteSQL = "Database.executeSQL"
	
	
	GetDatabaseTableNames = "Database.getDatabaseTableNames"
	
)

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// ExecuteSQL parameters
type ExecuteSQLParams struct {
	
	
	DatabaseId	DatabaseId	`json:"databaseId"`
	
	
	Query	string	`json:"query"`
	
}

// ExecuteSQL returns
type ExecuteSQLReturns struct {
	
	
	ColumnNames	[]string	`json:"columnNames"`
	
	
	Values	[]interface{}	`json:"values"`
	
	
	SqlError	Error	`json:"sqlError"`
	
}

// GetDatabaseTableNames parameters
type GetDatabaseTableNamesParams struct {
	
	
	DatabaseId	DatabaseId	`json:"databaseId"`
	
}

// GetDatabaseTableNames returns
type GetDatabaseTableNamesReturns struct {
	
	
	TableNames	[]string	`json:"tableNames"`
	
}

