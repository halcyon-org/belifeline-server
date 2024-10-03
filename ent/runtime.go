// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/halcyon-org/kizuna/ent/externalinformation"
	"github.com/halcyon-org/kizuna/ent/koyodata"
	"github.com/halcyon-org/kizuna/ent/koyoinformation"
	"github.com/halcyon-org/kizuna/ent/schema"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	externalinformationFields := schema.ExternalInformation{}.Fields()
	_ = externalinformationFields
	// externalinformationDescName is the schema descriptor for name field.
	externalinformationDescName := externalinformationFields[1].Descriptor()
	// externalinformation.NameValidator is a validator for the "name" field. It is called by the builders before save.
	externalinformation.NameValidator = externalinformationDescName.Validators[0].(func(string) error)
	// externalinformationDescDescription is the schema descriptor for description field.
	externalinformationDescDescription := externalinformationFields[2].Descriptor()
	// externalinformation.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	externalinformation.DescriptionValidator = externalinformationDescDescription.Validators[0].(func(string) error)
	// externalinformationDescLicense is the schema descriptor for license field.
	externalinformationDescLicense := externalinformationFields[3].Descriptor()
	// externalinformation.LicenseValidator is a validator for the "license" field. It is called by the builders before save.
	externalinformation.LicenseValidator = externalinformationDescLicense.Validators[0].(func(string) error)
	// externalinformationDescFirstEntryAt is the schema descriptor for first_entry_at field.
	externalinformationDescFirstEntryAt := externalinformationFields[5].Descriptor()
	// externalinformation.DefaultFirstEntryAt holds the default value on creation for the first_entry_at field.
	externalinformation.DefaultFirstEntryAt = externalinformationDescFirstEntryAt.Default.(func() time.Time)
	// externalinformationDescLastUpdatedAt is the schema descriptor for last_updated_at field.
	externalinformationDescLastUpdatedAt := externalinformationFields[6].Descriptor()
	// externalinformation.DefaultLastUpdatedAt holds the default value on creation for the last_updated_at field.
	externalinformation.DefaultLastUpdatedAt = externalinformationDescLastUpdatedAt.Default.(func() time.Time)
	// externalinformation.UpdateDefaultLastUpdatedAt holds the default value on update for the last_updated_at field.
	externalinformation.UpdateDefaultLastUpdatedAt = externalinformationDescLastUpdatedAt.UpdateDefault.(func() time.Time)
	// externalinformationDescID is the schema descriptor for id field.
	externalinformationDescID := externalinformationFields[0].Descriptor()
	// externalinformation.DefaultID holds the default value on creation for the id field.
	externalinformation.DefaultID = externalinformationDescID.Default.(func() pulid.ID)
	koyodataFields := schema.KoyoData{}.Fields()
	_ = koyodataFields
	// koyodataDescVersion is the schema descriptor for version field.
	koyodataDescVersion := koyodataFields[4].Descriptor()
	// koyodata.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	koyodata.VersionValidator = koyodataDescVersion.Validators[0].(func(string) error)
	// koyodataDescEntryAt is the schema descriptor for entry_at field.
	koyodataDescEntryAt := koyodataFields[5].Descriptor()
	// koyodata.DefaultEntryAt holds the default value on creation for the entry_at field.
	koyodata.DefaultEntryAt = koyodataDescEntryAt.Default.(func() time.Time)
	// koyodataDescID is the schema descriptor for id field.
	koyodataDescID := koyodataFields[0].Descriptor()
	// koyodata.DefaultID holds the default value on creation for the id field.
	koyodata.DefaultID = koyodataDescID.Default.(func() pulid.ID)
	koyoinformationFields := schema.KoyoInformation{}.Fields()
	_ = koyoinformationFields
	// koyoinformationDescName is the schema descriptor for name field.
	koyoinformationDescName := koyoinformationFields[1].Descriptor()
	// koyoinformation.NameValidator is a validator for the "name" field. It is called by the builders before save.
	koyoinformation.NameValidator = koyoinformationDescName.Validators[0].(func(string) error)
	// koyoinformationDescVersion is the schema descriptor for version field.
	koyoinformationDescVersion := koyoinformationFields[5].Descriptor()
	// koyoinformation.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	koyoinformation.VersionValidator = koyoinformationDescVersion.Validators[0].(func(string) error)
	// koyoinformationDescLicense is the schema descriptor for license field.
	koyoinformationDescLicense := koyoinformationFields[6].Descriptor()
	// koyoinformation.LicenseValidator is a validator for the "license" field. It is called by the builders before save.
	koyoinformation.LicenseValidator = koyoinformationDescLicense.Validators[0].(func(string) error)
	// koyoinformationDescFirstEntryAt is the schema descriptor for first_entry_at field.
	koyoinformationDescFirstEntryAt := koyoinformationFields[8].Descriptor()
	// koyoinformation.DefaultFirstEntryAt holds the default value on creation for the first_entry_at field.
	koyoinformation.DefaultFirstEntryAt = koyoinformationDescFirstEntryAt.Default.(func() time.Time)
	// koyoinformationDescUpdatedAt is the schema descriptor for updated_at field.
	koyoinformationDescUpdatedAt := koyoinformationFields[10].Descriptor()
	// koyoinformation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	koyoinformation.DefaultUpdatedAt = koyoinformationDescUpdatedAt.Default.(func() time.Time)
	// koyoinformation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	koyoinformation.UpdateDefaultUpdatedAt = koyoinformationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// koyoinformationDescID is the schema descriptor for id field.
	koyoinformationDescID := koyoinformationFields[0].Descriptor()
	// koyoinformation.DefaultID holds the default value on creation for the id field.
	koyoinformation.DefaultID = koyoinformationDescID.Default.(func() pulid.ID)
}
