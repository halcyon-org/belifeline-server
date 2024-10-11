// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminUsersColumns holds the columns for the "admin_users" table.
	AdminUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_used_at", Type: field.TypeTime},
		{Name: "last_updated_at", Type: field.TypeTime},
	}
	// AdminUsersTable holds the schema information for the "admin_users" table.
	AdminUsersTable = &schema.Table{
		Name:       "admin_users",
		Columns:    AdminUsersColumns,
		PrimaryKey: []*schema.Column{AdminUsersColumns[0]},
	}
	// ClientDataColumns holds the columns for the "client_data" table.
	ClientDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_used_at", Type: field.TypeTime},
		{Name: "last_updated_at", Type: field.TypeTime},
	}
	// ClientDataTable holds the schema information for the "client_data" table.
	ClientDataTable = &schema.Table{
		Name:       "client_data",
		Columns:    ClientDataColumns,
		PrimaryKey: []*schema.Column{ClientDataColumns[0]},
	}
	// ExternalInformationsColumns holds the columns for the "external_informations" table.
	ExternalInformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "license", Type: field.TypeString},
		{Name: "license_description", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString},
		{Name: "first_entry_at", Type: field.TypeTime},
		{Name: "last_updated_at", Type: field.TypeTime},
		{Name: "koyo_information_externals", Type: field.TypeString, Nullable: true},
	}
	// ExternalInformationsTable holds the schema information for the "external_informations" table.
	ExternalInformationsTable = &schema.Table{
		Name:       "external_informations",
		Columns:    ExternalInformationsColumns,
		PrimaryKey: []*schema.Column{ExternalInformationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "external_informations_koyo_informations_externals",
				Columns:    []*schema.Column{ExternalInformationsColumns[8]},
				RefColumns: []*schema.Column{KoyoInformationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// KoyoDataColumns holds the columns for the "koyo_data" table.
	KoyoDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "koyo_id", Type: field.TypeString, Unique: true},
		{Name: "scale", Type: field.TypeFloat64},
		{Name: "params", Type: field.TypeJSON},
		{Name: "version", Type: field.TypeString},
		{Name: "entry_at", Type: field.TypeTime},
		{Name: "target_at", Type: field.TypeTime},
		{Name: "koyo_information_data", Type: field.TypeString, Nullable: true},
	}
	// KoyoDataTable holds the schema information for the "koyo_data" table.
	KoyoDataTable = &schema.Table{
		Name:       "koyo_data",
		Columns:    KoyoDataColumns,
		PrimaryKey: []*schema.Column{KoyoDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "koyo_data_koyo_informations_data",
				Columns:    []*schema.Column{KoyoDataColumns[7]},
				RefColumns: []*schema.Column{KoyoInformationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// KoyoInformationsColumns holds the columns for the "koyo_informations" table.
	KoyoInformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "params", Type: field.TypeJSON},
		{Name: "scales", Type: field.TypeJSON},
		{Name: "version", Type: field.TypeString},
		{Name: "license", Type: field.TypeString},
		{Name: "data_type", Type: field.TypeEnum, Enums: []string{"unspecified", "image", "csv", "json"}},
		{Name: "api_key", Type: field.TypeString},
		{Name: "first_entry_at", Type: field.TypeTime},
		{Name: "last_entry_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// KoyoInformationsTable holds the schema information for the "koyo_informations" table.
	KoyoInformationsTable = &schema.Table{
		Name:       "koyo_informations",
		Columns:    KoyoInformationsColumns,
		PrimaryKey: []*schema.Column{KoyoInformationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminUsersTable,
		ClientDataTable,
		ExternalInformationsTable,
		KoyoDataTable,
		KoyoInformationsTable,
	}
)

func init() {
	ExternalInformationsTable.ForeignKeys[0].RefTable = KoyoInformationsTable
	KoyoDataTable.ForeignKeys[0].RefTable = KoyoInformationsTable
}
