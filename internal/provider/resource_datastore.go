package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceDatastoreType struct {
}

// GetSchema - Datastore Resource schema
func (r resourceDatastoreType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Single Datastore.",
		MarkdownDescription: "Single **Datastore**.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"resource_group_name": {
				Type:     types.StringType,
				Required: true,
			},
			"workspace_name": {
				Type:     types.StringType,
				Required: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"description": {
				Type:     types.StringType,
				Required: true,
			},
			"is_default": {
				Type:     types.BoolType,
				Optional: true,
			},
			"storage_type": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					NewStorageTypeValidator(),
				},
			},
			"storage_account_name": {
				Type:       types.StringType,
				Optional:   true,
				Validators: []tfsdk.AttributeValidator{
					//StorageAccountNameValidator{},
				},
			},
			"storage_container_name": {
				Type:     types.StringType,
				Optional: true,
			},
			"auth": {
				Required: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"credentials_type": {
						Type:     types.StringType,
						Required: true,
					},
					"client_id": {
						Type:     types.StringType,
						Optional: true,
					},
					"client_secret": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
					"account_key": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
					"sql_user_name": {
						Type:     types.StringType,
						Optional: true,
					},
					"sql_user_password": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
				}),
			},
			"system_data": {
				Computed: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"creation_date": {
						Type:     types.StringType,
						Computed: true,
					},
					"creation_user": {
						Type:     types.StringType,
						Computed: true,
					},
					"creation_user_type": {
						Type:     types.StringType,
						Computed: true,
					},
					"last_modified_date": {
						Type:     types.StringType,
						Computed: true,
					},
					"last_modified_user": {
						Type:     types.StringType,
						Computed: true,
					},
					"last_modified_user_type": {
						Type:     types.StringType,
						Computed: true,
					},
				}),
			},
		},
	}, nil
}

// NewResource  - New Datastore resource instance
func (r resourceDatastoreType) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceDatastore{
		p: *(p.(*provider)),
	}, nil
}

type resourceDatastore struct {
	p provider
}

func (r resourceDatastore) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	panic("implement me")
}

func (r resourceDatastore) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	panic("implement me")
}

func (r resourceDatastore) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	panic("implement me")
}

func (r resourceDatastore) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	panic("implement me")
}

func (r resourceDatastore) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
