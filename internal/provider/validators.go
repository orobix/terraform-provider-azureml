package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
	return fmt.Sprintf("Accepted values are: %v", s.allowedTypes)
}

func (s StorageTypeValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Accepted values are: %v", s.allowedTypes)
}

func (s StorageTypeValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	v, ok := request.AttributeConfig.(types.String)
	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"Attribute value should be a string",
		)
		return
	}
	if !contains(s.allowedTypes, v.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid storage type",
			fmt.Sprintf("Allowed storage types are: %v", s.allowedTypes),
		)
	}
}

type StorageAccountNameValidator struct {
}

func (s StorageAccountNameValidator) Description(ctx context.Context) string {
	panic("implement me")
}

func (s StorageAccountNameValidator) MarkdownDescription(ctx context.Context) string {
	panic("implement me")
}

func (s StorageAccountNameValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	panic("implement me")
}
