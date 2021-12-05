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
		Description: "Use this resource to access the information of a specific Datastore of a certain Azure ML " +
			"Workspace. Authentication credentials are not included in the provided information.",
		Attributes: map[string]tfsdk.Attribute{
			"resource_group_name": {
				Type:        types.StringType,
				Required:    true,
				Description: "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
			},
			"workspace_name": {
				Type:        types.StringType,
				Required:    true,
				Description: "The name of the Azure ML Workspace to which the datastore belongs to.",
			},
			"name": {
				Type:        types.StringType,
				Required:    true,
				Description: "The name of the datastore.",
			},
			"id": {
				Type:        types.StringType,
				Computed:    true,
				Description: "The ID of the datastore.",
			},
			"description": {
				Type:        types.StringType,
				Computed:    true,
				Description: "The description of the datastore.",
			},
			"is_default": {
				Type:        types.BoolType,
				Computed:    true,
				Description: "Is the datastore the default datastore of the Azure ML Workspace?",
			},
			"storage_type": {
				Type:     types.StringType,
				Computed: true,
				Description: fmt.Sprintf(
					"The type of the storage to which the datstore is linked to. Possible values are: %v",
					NewStorageTypeValidator().allowedTypes,
				),
			},
			"storage_account_name": {
				Type:        types.StringType,
				Computed:    true,
				Description: "The name of the Storage Account to which the datastore is linked to.",
			},
			"storage_container_name": {
				Type:        types.StringType,
				Computed:    true,
				Description: "The name of the Storage Container to which the datastore is linked to.",
			},
			"credentials_type": {
				Type:     types.StringType,
				Computed: true,
				Description: fmt.Sprintf(
					"The type of credentials used for authenticating with the underlying storage. "+
						"Possible values are: %v.",
					NewDatastoreCredentialsTypeValidator().allowedTypes,
				),
			},
			"system_data": {
				Computed: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"creation_date": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The timestamp corresponding to the creation of the datastore.",
					},
					"creation_user": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The user that created the datastore.",
					},
					"creation_user_type": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The kind of user that created the datastore (Service Principal or User).",
					},
					"last_modified_date": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The timestamp corresponding to the last update of the datastore.",
					},
					"last_modified_user": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The user that last updated the datastore.",
					},
					"last_modified_user_type": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The kind of user that last updated the datastore (Service Principal or User).",
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
	var resourceData ConfigReadableDatastore

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
		msg := fmt.Sprintf("Error retrieving datastore \"%s\".", resourceData.Name.Value)
		resp.Diagnostics.AddError(msg, err.Error())
		return
	}

	// Update resource data with fetched data
	resourceData.ID = types.String{Value: datastore.Id}
	resourceData.Name = types.String{Value: datastore.Name}
	resourceData.Description = types.String{Value: datastore.Description}
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
