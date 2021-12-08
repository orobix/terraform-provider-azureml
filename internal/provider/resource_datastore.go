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
		Description: "Manages a Datastore of an Azure Machine Learning workspace.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:        types.StringType,
				Computed:    true,
				Description: "The ID of the datastore.",
			},
			"resource_group_name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					StringNotEmptyValidator{},
				},
				Description: "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
			},
			"workspace_name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					StringNotEmptyValidator{},
				},
				Description: "The name of the Azure ML Workspace to which the datastore belongs to.",
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Validators: []tfsdk.AttributeValidator{
					StringNotEmptyValidator{},
				},
				Description: "The name of the datastore.",
			},
			"description": {
				Type:        types.StringType,
				Required:    true,
				Description: "The description of the datastore.",
			},
			"is_default": {
				Type:        types.BoolType,
				Optional:    true,
				Computed:    true,
				Description: "Is the datastore the default datastore of the Azure ML Workspace?",
			},
			"storage_type": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					NewStorageTypeValidator(),
				},
				Description: fmt.Sprintf(
					"The type of the storage to which the datstore is linked to. Possible values are: %v",
					NewStorageTypeValidator().allowedTypes,
				),
			},
			"storage_account_name": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					StringNotEmptyValidator{},
					StorageAccountNameValidator{},
				},
				Description: "The name of the Storage Account to which the datastore is linked to.",
			},
			"storage_container_name": {
				Type:     types.StringType,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					StringNotEmptyValidator{},
				},
				Description: "The name of the Storage Container to which the datastore is linked to.",
			},
			"auth": {
				Required: true,
				Description: "The credentials for authenticating with the storage linked to the datastore. " +
					"The authentication methods depends on the underlying storage type of the datastore.",
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"credentials_type": {
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							NewDatastoreCredentialsTypeValidator(),
						},
						Description: fmt.Sprintf(
							"The type of credentials used for authenticating with the underlying storage. "+
								"Possible values are: %v.",
							NewDatastoreCredentialsTypeValidator().allowedTypes,
						),
					},
					"tenant_id": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The ID of the tenant to which the Service Principal used for authenticating " +
							"belongs to.",
					},
					"client_id": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The application ID of the service principal used for authenticating with the " +
							"underlying storage of the datastore.",
					},
					"client_secret": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The client secret of the service principal used for authenticating with the " +
							"underlying storage of the datastore.",
					},
					"account_key": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The primary key of the Storage Account linked to the datastore.",
					},
					"sql_user_name": {
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The username of the identity used for authenticating with the SQL database linked " +
							"to the storage account.",
					},
					"sql_user_password": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
						Validators: []tfsdk.AttributeValidator{
							StringNotEmptyValidator{},
						},
						Description: "The password of the identity used for authenticating with the SQL database linked " +
							"to the storage account.",
					},
				}),
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
	var resourceData ConfigReadableDatastoreWithAuth

	diags := req.Config.Get(ctx, &resourceData)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newDatastore := workspace.Datastore{
		Name:                 resourceData.Name.Value,
		IsDefault:            resourceData.IsDefault.Value,
		Description:          resourceData.Description.Value,
		StorageType:          resourceData.StorageType.Value,
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
		resp.Diagnostics.AddError("Error creating datastore.", err.Error())
		return
	}

	auth := DatastoreAuth{
		CredentialsType: types.String{Value: createdDatastore.Auth.CredentialsType},
		// read from resource data since APIs do not return secrets
		ClientSecret:    resourceData.Auth.ClientSecret,
		AccountKey:      resourceData.Auth.AccountKey,
		SqlUserPassword: resourceData.Auth.SqlUserPassword,
	}
	if resourceData.Auth.ClientId.Null == true {
		auth.ClientId = types.String{Null: true}
	} else {
		auth.ClientId = types.String{Value: createdDatastore.Auth.ClientId}
	}
	if resourceData.Auth.TenantId.Null == true {
		auth.TenantId = types.String{Null: true}
	} else {
		auth.TenantId = types.String{Value: createdDatastore.Auth.TenantId}
	}
	if resourceData.Auth.SqlUserName.Null == true {
		auth.SqlUserName = types.String{Null: true}
	} else {
		auth.SqlUserName = types.String{Value: createdDatastore.Auth.SqlUserName}
	}

	state := DatastoreWithAuth{
		ResourceGroupName:    resourceData.ResourceGroupName,
		WorkspaceName:        resourceData.WorkspaceName,
		ID:                   types.String{Value: createdDatastore.Id},
		Name:                 types.String{Value: createdDatastore.Name},
		Description:          types.String{Value: createdDatastore.Description},
		IsDefault:            types.Bool{Value: createdDatastore.IsDefault},
		StorageType:          types.String{Value: createdDatastore.StorageType},
		StorageAccountName:   types.String{Value: createdDatastore.StorageAccountName},
		StorageContainerName: types.String{Value: createdDatastore.StorageContainerName},
		Auth:                 auth,
		SystemData: SystemData{
			CreationDate:         types.String{Value: createdDatastore.SystemData.CreationDate.Format(defaultDateFormat)},
			CreationUser:         types.String{Value: createdDatastore.SystemData.CreationUser},
			CreationUserType:     types.String{Value: createdDatastore.SystemData.CreationUserType},
			LastModifiedDate:     types.String{Value: createdDatastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
			LastModifiedUser:     types.String{Value: createdDatastore.SystemData.LastModifiedUser},
			LastModifiedUserType: types.String{Value: createdDatastore.SystemData.LastModifiedUserType},
		},
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceDatastore) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state DatastoreWithAuth

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	datastore, err := r.p.client.GetDatastore(
		state.ResourceGroupName.Value,
		state.WorkspaceName.Value,
		state.Name.Value,
	)
	if err != nil {
		msg := fmt.Sprintf("Error retrieving datastore \"%s\"", state.Name.Value)
		resp.Diagnostics.AddError(msg, err.Error())
		return
	}

	// Update state with fetched data
	state = mergeToDatastoreWithAuth(&state, datastore)

	// Set entire state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceDatastore) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Retrieve the changes proposed in the execution plan
	var plan ConfigReadableDatastoreWithAuth
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve the current state values
	var state DatastoreWithAuth
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update the resource on AML
	patchDatastore := workspace.Datastore{
		Id:                   state.ID.Value,
		Name:                 plan.Name.Value,
		IsDefault:            plan.IsDefault.Value,
		Description:          plan.Description.Value,
		StorageType:          plan.StorageType.Value,
		StorageAccountName:   plan.StorageAccountName.Value,
		StorageContainerName: plan.StorageContainerName.Value,
		Auth: workspace.DatastoreAuth{
			CredentialsType: plan.Auth.CredentialsType.Value,
			ClientId:        plan.Auth.CredentialsType.Value,
			TenantId:        plan.Auth.TenantId.Value,
			ClientSecret:    plan.Auth.ClientSecret.Value,
			AccountKey:      plan.Auth.AccountKey.Value,
			SqlUserName:     plan.Auth.SqlUserName.Value,
			SqlUserPassword: plan.Auth.SqlUserPassword.Value,
		},
	}
	updatedDatastore, err := r.p.client.CreateOrUpdateDatastore(
		plan.ResourceGroupName.Value,
		plan.WorkspaceName.Value,
		&patchDatastore,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error updating datastore \"%s\".", state.Name.Value),
			err.Error(),
		)
		return
	}

	// Write to state the updated resource
	newState := mergeToDatastoreWithAuth(&state, updatedDatastore)
	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
}

func (r resourceDatastore) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	panic("implement me")
}

func (r resourceDatastore) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	panic("implement me")
}
