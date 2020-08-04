package database

import "fmt"

const resourceInUseException = "ResourceInUseException: Cannot create preexisting table"
const resourceInUseExceptionPriceHistory = "ResourceInUseException: Table already exists: PriceHistory"

// getTableExistsExceptionString returns a table specific exeption string
func getTableExistsExceptionString(tableName string) string {
	return fmt.Sprintf("ResourceInUseException: Table already exists: %s", tableName)
}
