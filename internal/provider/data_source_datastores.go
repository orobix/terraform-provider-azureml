package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceDatastoresType struct{}

func (d dataSourceDatastoresType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "List of Datastores.",
		MarkdownDescription: "List of **Datastores**.",
		Attributes: map[string]tfsdk.Attribute{
			"resource_group_name": {
				Type:     types.StringType,
				Required: true,
			},
			"workspace_name": {
				Type:     types.StringType,
				Required: true,
			},
			"datastores": {
				Computed: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
					"resource_group_name": {
						Type:     types.StringType,
						Required: true,
					},
					"workspace_name": {
						Type:     types.StringType,
						Required: true,
					},
					"id": {
						Type:     types.StringType,
						Computed: true,
					},
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"description": {
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
		resp.Diagnostics.AddError("Error retrieving datastores", err.Error())
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
