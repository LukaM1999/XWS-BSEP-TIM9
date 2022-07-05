// Code generated by entc, DO NOT EDIT.

package skill

const (
	// Label holds the string label denoting the skill type in the database.
	Label = "skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeRequired holds the string denoting the required edge name in mutations.
	EdgeRequired = "required"
	// Table holds the table name of the skill in the database.
	Table = "skills"
	// RequiredTable is the table that holds the required relation/edge. The primary key declared below.
	RequiredTable = "job_offer_requires"
	// RequiredInverseTable is the table name for the JobOffer entity.
	// It exists in this package in order to avoid circular dependency with the "joboffer" package.
	RequiredInverseTable = "job_offers"
)

// Columns holds all SQL columns for skill fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// RequiredPrimaryKey and RequiredColumn2 are the table columns denoting the
	// primary key for the required relation (M2M).
	RequiredPrimaryKey = []string{"job_offer_id", "skill_id"}
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
