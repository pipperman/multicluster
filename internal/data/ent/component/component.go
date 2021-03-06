// Code generated by entc, DO NOT EDIT.

package component

const (
	// Label holds the string label denoting the component type in the database.
	Label = "component"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the component in the database.
	Table = "components"
)

// Columns holds all SQL columns for component fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
