package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceDatastoresType struct{}

func (d dataSourceDatastoresType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Use this resource to fetch the list of of the Datastores belonging to a certain " +
			"Azure ML Workspace. Authentication credentials are not included in the returned datastore list.",
		Attributes: map[string]tfsdk.Attribute{
			"resource_group_name": {
				Type:        types.StringType,
				Required:    true,
				Description: "The name of the resource group of the Azure ML Workspace.",
			},
			"workspace_name": {
				Type:        types.StringType,
				Required:    true,
				Description: "The name of the Azure ML Workspace from which the datastores will be fetched.",
			},
			"datastores": {
				Computed: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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
					"id": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The ID of the datastore.",
					},
					"name": {
						Type:        types.StringType,
						Computed:    true,
						Description: "The name of the datastore.",
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
				}, tfsdk.SetNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (d dataSourceDatastoresType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceDatastores{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceDatastores struct {
	p provider
}

func (d dataSourceDatastores) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var resourceData DatastoreList

	diags := req.Config.Get(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	datastores, err := d.p.client.GetDatastores(resourceData.ResourceGroupName.Value, resourceData.WorkspaceName.Value)
	if err != nil {
		resp.Diagnostics.AddError("Error retrieving datastores.", err.Error())
		return
	}

	// Map response body to resource schema
	for _, datastore := range datastores {
		d := Datastore{
			ResourceGroupName: types.String{Value: resourceData.ResourceGroupName.Value},
			WorkspaceName:     types.String{Value: resourceData.WorkspaceName.Value},

			ID:          types.String{Value: datastore.Id},
			Name:        types.String{Value: datastore.Name},
			Description: types.String{Value: datastore.Description},
			IsDefault:   types.Bool{Value: datastore.IsDefault},

			StorageType:          types.String{Value: datastore.StorageType},
			StorageAccountName:   types.String{Value: datastore.StorageAccountName},
			StorageContainerName: types.String{Value: datastore.StorageContainerName},

			CredentialsType: types.String{Value: datastore.Auth.CredentialsType},
			SystemData: SystemData{
				CreationDate:         types.String{Value: datastore.SystemData.CreationDate.Format(defaultDateFormat)},
				CreationUser:         types.String{Value: datastore.SystemData.CreationUser},
				CreationUserType:     types.String{Value: datastore.SystemData.CreationUserType},
				LastModifiedDate:     types.String{Value: datastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
				LastModifiedUser:     types.String{Value: datastore.SystemData.LastModifiedUser},
				LastModifiedUserType: types.String{Value: datastore.SystemData.LastModifiedUserType},
			},
		}
		resourceData.Datastores = append(resourceData.Datastores, d)
	}

	// Set state
	diags = resp.State.Set(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
