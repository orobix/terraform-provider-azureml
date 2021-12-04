package provider

import (
	"fmt"
	"github.com/Telemaco019/azureml-go-sdk/workspace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

import (
	"context"
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
					"tenant_id": {
						Type:     types.StringType,
						Optional: true,
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

func (r resourceDatastore) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var resourceData WriteDatastoreWithSystemDataObject

	diags := req.Config.Get(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newDatastore := workspace.Datastore{
		Name:                 resourceData.Name.Value,
		IsDefault:            resourceData.IsDefault.Value,
		Description:          resourceData.Description.Value,
		Type:                 resourceData.StorageType.Value,
		StorageAccountName:   resourceData.StorageAccountName.Value,
		StorageContainerName: resourceData.StorageContainerName.Value,
		Auth: workspace.DatastoreAuth{
			CredentialsType: resourceData.Auth.CredentialsType.Value,
			ClientId:        resourceData.Auth.ClientId.Value,
			TenantId:        resourceData.Auth.TenantId.Value,
			ClientSecret:    resourceData.Auth.ClientSecret.Value,
			AccountKey:      resourceData.Auth.AccountKey.Value,
			SqlUserName:     resourceData.Auth.SqlUserName.Value,
			SqlUserPassword: resourceData.Auth.SqlUserPassword.Value,
		},
	}
	createdDatastore, err := r.p.client.CreateOrUpdateDatastore(
		resourceData.ResourceGroupName.Value,
		resourceData.WorkspaceName.Value,
		&newDatastore,
	)
	if err != nil {
		resp.Diagnostics.AddError("Error creating datastore", err.Error())
	}

	result := WriteDatastoreWithSystemDataStruct{
		ResourceGroupName:    types.String{Value: resourceData.ResourceGroupName.Value},
		WorkspaceName:        types.String{Value: resourceData.WorkspaceName.Value},
		ID:                   types.String{Value: createdDatastore.Id},
		Name:                 types.String{Value: createdDatastore.Name},
		Description:          types.String{Value: createdDatastore.Description},
		IsDefault:            types.Bool{Value: createdDatastore.IsDefault},
		StorageType:          types.String{Value: createdDatastore.StorageType},
		StorageAccountName:   types.String{Value: createdDatastore.StorageAccountName},
		StorageContainerName: types.String{Value: createdDatastore.StorageContainerName},
		Auth: DatastoreAuth{
			CredentialsType: types.String{Value: createdDatastore.Auth.CredentialsType},
			TenantId:        types.String{Value: createdDatastore.Auth.TenantId},
			ClientId:        types.String{Value: createdDatastore.Auth.ClientId},
			SqlUserName:     types.String{Value: createdDatastore.Auth.SqlUserName},
			// read from resource data since APIs do not return secrets
			ClientSecret:    types.String{Value: resourceData.Auth.ClientSecret.Value},
			AccountKey:      types.String{Value: resourceData.Auth.AccountKey.Value},
			SqlUserPassword: types.String{Value: resourceData.Auth.SqlUserPassword.Value},
		},
		SystemData: SystemData{
			CreationDate:         types.String{Value: createdDatastore.SystemData.CreationDate.Format(defaultDateFormat)},
			CreationUser:         types.String{Value: createdDatastore.SystemData.CreationUser},
			CreationUserType:     types.String{Value: createdDatastore.SystemData.CreationUserType},
			LastModifiedDate:     types.String{Value: createdDatastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
			LastModifiedUser:     types.String{Value: createdDatastore.SystemData.LastModifiedUser},
			LastModifiedUserType: types.String{Value: createdDatastore.SystemData.LastModifiedUserType},
		},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceDatastore) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var resourceData WriteDatastoreWithSystemDataStruct

	diags := req.State.Get(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	datastore, err := r.p.client.GetDatastore(
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
	resourceData.IsDefault = types.Bool{Value: datastore.IsDefault}
	resourceData.StorageType = types.String{Value: datastore.StorageType}
	resourceData.StorageAccountName = types.String{Value: datastore.StorageAccountName}
	resourceData.StorageContainerName = types.String{Value: datastore.StorageContainerName}
	resourceData.Auth = DatastoreAuth{
		CredentialsType: types.String{Value: datastore.Auth.CredentialsType},
		TenantId:        types.String{Value: datastore.Auth.TenantId},
		ClientId:        types.String{Value: datastore.Auth.ClientId},
		SqlUserName:     types.String{Value: datastore.Auth.SqlUserName},
		// Use resource data values since APIs do not return secrets
		ClientSecret:    types.String{Value: resourceData.Auth.ClientSecret.Value},
		AccountKey:      types.String{Value: resourceData.Auth.AccountKey.Value},
		SqlUserPassword: types.String{Value: resourceData.Auth.SqlUserPassword.Value},
	}
	resourceData.SystemData = SystemData{
		CreationDate:         types.String{Value: datastore.SystemData.CreationDate.Format(defaultDateFormat)},
		CreationUser:         types.String{Value: datastore.SystemData.CreationUser},
		CreationUserType:     types.String{Value: datastore.SystemData.CreationUserType},
		LastModifiedDate:     types.String{Value: datastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
		LastModifiedUser:     types.String{Value: datastore.SystemData.LastModifiedUser},
		LastModifiedUserType: types.String{Value: datastore.SystemData.LastModifiedUserType},
	}

	// Set entire state
	diags = resp.State.Set(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
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
