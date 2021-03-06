// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgePet holds the string denoting the pet edge name in mutations.
	EdgePet = "pet"

	// Table holds the table name of the user in the database.
	Table = "users"
	// PetTable is the table the holds the pet relation/edge. The primary key declared below.
	PetTable = "user_pet"
	// PetInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetInverseTable = "pets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// PetPrimaryKey and PetColumn2 are the table columns denoting the
	// primary key for the pet relation (M2M).
	PetPrimaryKey = []string{"user_id", "pet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
