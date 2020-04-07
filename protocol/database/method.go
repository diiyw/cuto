package database

// Disables database tracking, prevents database events from being sent to the client.
const Disable = "Database.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables database tracking, database events will now be delivered to the client.
const Enable = "Database.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// 
const ExecuteSQL = "Database.executeSQL"

type ExecuteSQLParams struct {

	// 
	DatabaseId 	DatabaseId	`json:"databaseId"`

	// 
	Query 	string	`json:"query"`
}

type ExecuteSQLResult struct {

	// 
	ColumnNames 	[]string	`json:"columnNames"`
	// 
	Values 	[]interface{}	`json:"values"`
	// 
	SqlError 	Error	`json:"sqlError"`
}

// 
const GetDatabaseTableNames = "Database.getDatabaseTableNames"

type GetDatabaseTableNamesParams struct {

	// 
	DatabaseId 	DatabaseId	`json:"databaseId"`
}

type GetDatabaseTableNamesResult struct {

	// 
	TableNames 	[]string	`json:"tableNames"`
}