// Code generated by ent, DO NOT EDIT.

package koyodata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the koyodata type in the database.
	Label = "koyo_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKoyoID holds the string denoting the koyo_id field in the database.
	FieldKoyoID = "koyo_id"
	// FieldScale holds the string denoting the scale field in the database.
	FieldScale = "scale"
	// FieldParams holds the string denoting the params field in the database.
	FieldParams = "params"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldEntryAt holds the string denoting the entry_at field in the database.
	FieldEntryAt = "entry_at"
	// FieldTargetAt holds the string denoting the target_at field in the database.
	FieldTargetAt = "target_at"
	// Table holds the table name of the koyodata in the database.
	Table = "koyo_data"
)

// Columns holds all SQL columns for koyodata fields.
var Columns = []string{
	FieldID,
	FieldKoyoID,
	FieldScale,
	FieldParams,
	FieldVersion,
	FieldEntryAt,
	FieldTargetAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "koyo_data"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"koyo_information_data",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// VersionValidator is a validator for the "version" field. It is called by the builders before save.
	VersionValidator func(string) error
	// DefaultEntryAt holds the default value on creation for the "entry_at" field.
	DefaultEntryAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the KoyoData queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKoyoID orders the results by the koyo_id field.
func ByKoyoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKoyoID, opts...).ToFunc()
}

// ByScale orders the results by the scale field.
func ByScale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScale, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByEntryAt orders the results by the entry_at field.
func ByEntryAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryAt, opts...).ToFunc()
}

// ByTargetAt orders the results by the target_at field.
func ByTargetAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetAt, opts...).ToFunc()
}
