package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	storageAccountNameMaxLength = 24
	storageAccountNameMinLength = 24
)

func NewDatastoreCredentialsTypeValidator() *DatastoreCredentialsTypeValidator {
	return &DatastoreCredentialsTypeValidator{
		allowedTypes: []string{
			"AccountKey",
			"Certificate",
			"None",
			"Sas",
			"ServicePrincipal",
			"SqlAdmin",
		},
	}
}

type DatastoreCredentialsTypeValidator struct {
	allowedTypes []string
}

func (d DatastoreCredentialsTypeValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Accepted values are: %v.", d.allowedTypes)
}

func (d DatastoreCredentialsTypeValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Accepted values are: %v.", d.allowedTypes)
}

func (d DatastoreCredentialsTypeValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	v, ok := request.AttributeConfig.(types.String)
	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			"Attribute value should be a string.",
		)
		return
	}
	if !contains(d.allowedTypes, v.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid datastore credential type.",
			fmt.Sprintf("Allowed credentials types are: %v.", d.allowedTypes),
		)
	}
}

func NewStorageTypeValidator() *StorageTypeValidator {
	return &StorageTypeValidator{
		allowedTypes: []string{
			"AzureFile",
			"AzureBlob",
			"AzureDataLakeGen1",
			"AzureDataLakeGen2",
			"AzureMySql",
			"AzurePostgreSql",
			"AzureSqlDatabase",
			"GlusterFs",
		},
	}
}

type StorageTypeValidator struct {
	allowedTypes []string
}

func (s StorageTypeValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Accepted values are: %v.", s.allowedTypes)
}

func (s StorageTypeValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Accepted values are: %v.", s.allowedTypes)
}

func (s StorageTypeValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	v, ok := request.AttributeConfig.(types.String)
	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			"Attribute value should be a string.",
		)
		return
	}
	if !contains(s.allowedTypes, v.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid storage type.",
			fmt.Sprintf("Allowed storage types are: %v.", s.allowedTypes),
		)
	}
}

type StorageAccountNameValidator struct {
	AttributeIsRequired bool
}

func (s StorageAccountNameValidator) Description(ctx context.Context) string {
	return "The attribute must be a valid name for a Storage Account."
}

func (s StorageAccountNameValidator) MarkdownDescription(ctx context.Context) string {
	return "The attribute must be a valid name for a Storage Account."
}

func (s StorageAccountNameValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	v, ok := request.AttributeConfig.(types.String)
	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"Attribute value should be a string",
		)
		return
	}
	if s.AttributeIsRequired == false && (v.Unknown == true || v.Null == true) {
		return
	}

	// Check length
	length := len(v.Value)
	if length < storageAccountNameMinLength || length > storageAccountNameMaxLength {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			fmt.Sprintf(
				"Storage account name must be between %d and %d characters.",
				storageAccountNameMinLength,
				storageAccountNameMaxLength,
			),
		)
		return
	}

	// Check format
	if !stringIsOnlyLettersAndDigits(v.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			"Storage account name can contain only characters and digits.",
		)
		return
	}
}

type StringNotEmptyValidator struct {
	AttributeIsRequired bool
}

func (s StringNotEmptyValidator) Description(ctx context.Context) string {
	return "The attribute must be a non-empty string."
}

func (s StringNotEmptyValidator) MarkdownDescription(ctx context.Context) string {
	return "The attribute must be a non-empty string."
}

func (s StringNotEmptyValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	v, ok := request.AttributeConfig.(types.String)
	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			"Attribute value should be a string.",
		)
		return
	}
	if s.AttributeIsRequired == false && (v.Unknown == true || v.Null == true) {
		return
	}
	if stringIsEmpty(v.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value.",
			"The value must be a non-empty string.",
		)
		return
	}
}
