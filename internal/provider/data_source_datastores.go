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
			"datastores": {
				Computed: true,
				Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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
					"auth": {
						Computed: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"credentials_type": {
								Type:     types.StringType,
								Computed: true,
							},
							"client_id": {
								Type:     types.StringType,
								Computed: true,
							},
							"client_secret": {
								Type:      types.StringType,
								Computed:  true,
								Sensitive: true,
							},
							"account_key": {
								Type:     types.StringType,
								Computed: true,
							},
							"sql_user_name": {
								Type:     types.StringType,
								Computed: true,
							},
							"sql_user_password": {
								Type:      types.StringType,
								Computed:  true,
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
	var resourceState struct {
		Datastores []Datastore `tfsdk:"datastores"`
	}

	datastores, err := d.p.client.GetDatastores()
	if err != nil {
		resp.Diagnostics.AddError("Error retrieving datastores", err.Error())
		return
	}

	// Map response body to resource schema
	for _, datastore := range datastores {
		d := Datastore{
			ID:          types.String{Value: datastore.Id},
			Name:        types.String{Value: datastore.Name},
			Description: types.String{Value: datastore.Description},
			Type:        types.String{Value: datastore.Type},
			IsDefault:   types.Bool{Value: datastore.IsDefault},

			StorageType:          types.String{Value: datastore.StorageType},
			StorageAccountName:   types.String{Value: datastore.StorageAccountName},
			StorageContainerName: types.String{Value: datastore.StorageContainerName},

			Auth: DatastoreAuth{
				CredentialsType: types.String{Value: datastore.Auth.CredentialsType},
				ClientId:        types.String{Value: datastore.Auth.ClientId},
				ClientSecret:    types.String{Value: datastore.Auth.ClientSecret},
				AccountKey:      types.String{Value: datastore.Auth.AccountKey},
				SqlUserName:     types.String{Value: datastore.Auth.SqlUserName},
				SqlUserPassword: types.String{Value: datastore.Auth.SqlUserPassword},
			},
			SystemData: SystemData{
				CreationDate:         types.String{Value: datastore.SystemData.CreationDate.Format(defaultDateFormat)},
				CreationUser:         types.String{Value: datastore.SystemData.CreationUser},
				CreationUserType:     types.String{Value: datastore.SystemData.CreationUserType},
				LastModifiedDate:     types.String{Value: datastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
				LastModifiedUser:     types.String{Value: datastore.SystemData.LastModifiedUser},
				LastModifiedUserType: types.String{Value: datastore.SystemData.LastModifiedUserType},
			},
		}

		resourceState.Datastores = append(resourceState.Datastores, d)
	}

	// Set state
	diags := resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
