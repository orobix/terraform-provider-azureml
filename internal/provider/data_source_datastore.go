package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type dataSourceDatastoreType struct{}

func (d dataSourceDatastoreType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Single ReadDatastoreWithSystemDataObject.",
		MarkdownDescription: "Single **ReadDatastoreWithSystemDataObject**.",
		Attributes: map[string]tfsdk.Attribute{
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
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"description": {
				Type:     types.StringType,
				Computed: true,
			},
			"type": {
				Type:     types.StringType,
				Computed: true,
			},
			"is_default": {
				Type:     types.BoolType,
				Computed: true,
			},
			"storage_type": {
				Type:     types.StringType,
				Computed: true,
			},
			"storage_account_name": {
				Type:     types.StringType,
				Computed: true,
			},
			"storage_container_name": {
				Type:     types.StringType,
				Computed: true,
			},
			"credentials_type": {
				Type:     types.StringType,
				Computed: true,
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

func (d dataSourceDatastoreType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceDatastore{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceDatastore struct {
	p provider
}

func (ds dataSourceDatastore) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var resourceData ReadDatastoreWithSystemDataObject

	diags := req.Config.Get(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	datastore, err := ds.p.client.GetDatastore(
		resourceData.ResourceGroupName.Value,
		resourceData.WorkspaceName.Value,
		resourceData.Name.Value,
	)
	if err != nil {
		msg := fmt.Sprintf("Error retrieving datastore \"%s\"", resourceData.Name.Value)
		resp.Diagnostics.AddError(msg, err.Error())
		return
	}

	// Update resource data with fetched data
	resourceData.ID = types.String{Value: datastore.Id}
	resourceData.Name = types.String{Value: datastore.Name}
	resourceData.Description = types.String{Value: datastore.Description}
	resourceData.Type = types.String{Value: datastore.Type}
	resourceData.IsDefault = types.Bool{Value: datastore.IsDefault}
	resourceData.StorageType = types.String{Value: datastore.StorageType}
	resourceData.StorageAccountName = types.String{Value: datastore.StorageAccountName}
	resourceData.StorageContainerName = types.String{Value: datastore.StorageContainerName}
	resourceData.CredentialsType = types.String{Value: datastore.Auth.CredentialsType}

	// Set entire state
	diags = resp.State.Set(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update state with object attributes
	systemData := SystemData{
		CreationDate:         types.String{Value: datastore.SystemData.CreationDate.Format(defaultDateFormat)},
		CreationUser:         types.String{Value: datastore.SystemData.CreationUser},
		CreationUserType:     types.String{Value: datastore.SystemData.CreationUserType},
		LastModifiedDate:     types.String{Value: datastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
		LastModifiedUser:     types.String{Value: datastore.SystemData.LastModifiedUser},
		LastModifiedUserType: types.String{Value: datastore.SystemData.LastModifiedUserType},
	}
	resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("system_data"), systemData)
}
