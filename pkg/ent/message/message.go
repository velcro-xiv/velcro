// Code generated by ent, DO NOT EDIT.

package message

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldSourceAddress holds the string denoting the source_address field in the database.
	FieldSourceAddress = "source_address"
	// FieldSourcePort holds the string denoting the source_port field in the database.
	FieldSourcePort = "source_port"
	// FieldDestinationAddress holds the string denoting the destination_address field in the database.
	FieldDestinationAddress = "destination_address"
	// FieldDestinationPort holds the string denoting the destination_port field in the database.
	FieldDestinationPort = "destination_port"
	// Table holds the table name of the message in the database.
	Table = "messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldTimestamp,
	FieldVersion,
	FieldSourceAddress,
	FieldSourcePort,
	FieldDestinationAddress,
	FieldDestinationPort,
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
