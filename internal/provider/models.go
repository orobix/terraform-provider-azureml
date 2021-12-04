package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemData struct {
	CreationDate     types.String `tfsdk:"creation_date"`
	CreationUser     types.String `tfsdk:"creation_user"`
	CreationUserType types.String `tfsdk:"creation_user_type"`

	LastModifiedDate     types.String `tfsdk:"last_modified_date"`
	LastModifiedUser     types.String `tfsdk:"last_modified_user"`
	LastModifiedUserType types.String `tfsdk:"last_modified_user_type"`
}

type DatastoreAuth struct {
	CredentialsType types.String `tfsdk:"credentials_type"`
	TenantId        types.String `tfsdk:"tenant_id"`
	ClientId        types.String `tfsdk:"client_id"`
	ClientSecret    types.String `tfsdk:"client_secret"`
	AccountKey      types.String `tfsdk:"account_key"`
	SqlUserName     types.String `tfsdk:"sql_user_name"`
	SqlUserPassword types.String `tfsdk:"sql_user_password"`
}

type Datastore struct {
	ResourceGroupName types.String `tfsdk:"resource_group_name"`
	WorkspaceName     types.String `tfsdk:"workspace_name"`

	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	StorageType          types.String `tfsdk:"storage_type"`
	StorageAccountName   types.String `tfsdk:"storage_account_name"`
	StorageContainerName types.String `tfsdk:"storage_container_name"`
	CredentialsType      types.String `tfsdk:"credentials_type"`

	SystemData SystemData `tfsdk:"system_data"`
}

// ConfigReadableDatastore - Datastore model that can be read from a Terraform config provided by the user, in which
// the SystemData can be Unknown, and it thus couldn't be written to the respective struct (SystemData field must
// therefore be of types.Object type).
type ConfigReadableDatastore struct {
	ResourceGroupName types.String `tfsdk:"resource_group_name"`
	WorkspaceName     types.String `tfsdk:"workspace_name"`

	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	StorageType          types.String `tfsdk:"storage_type"`
	StorageAccountName   types.String `tfsdk:"storage_account_name"`
	StorageContainerName types.String `tfsdk:"storage_container_name"`
	CredentialsType      types.String `tfsdk:"credentials_type"`

	SystemData types.Object `tfsdk:"system_data"`
}

type DatastoreList struct {
	ResourceGroupName types.String `tfsdk:"resource_group_name"`
	WorkspaceName     types.String `tfsdk:"workspace_name"`
	Datastores        []Datastore  `tfsdk:"datastores"`
}

type DatastoreWithAuth struct {
	ResourceGroupName types.String `tfsdk:"resource_group_name"`
	WorkspaceName     types.String `tfsdk:"workspace_name"`

	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	StorageType          types.String `tfsdk:"storage_type"`
	StorageAccountName   types.String `tfsdk:"storage_account_name"`
	StorageContainerName types.String `tfsdk:"storage_container_name"`

	Auth       DatastoreAuth `tfsdk:"auth"`
	SystemData SystemData    `tfsdk:"system_data"`
}

// ConfigReadableDatastoreWithAuth - Datastore model (including auth information) that can be read from a Terraform
// config provided by the user, in which the SystemData can be Unknown, and it thus couldn't be written to the
// respective struct (SystemData field must therefore be of types.Object type).
type ConfigReadableDatastoreWithAuth struct {
	ResourceGroupName types.String `tfsdk:"resource_group_name"`
	WorkspaceName     types.String `tfsdk:"workspace_name"`

	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	StorageType          types.String `tfsdk:"storage_type"`
	StorageAccountName   types.String `tfsdk:"storage_account_name"`
	StorageContainerName types.String `tfsdk:"storage_container_name"`

	Auth       DatastoreAuth `tfsdk:"auth"`
	SystemData types.Object  `tfsdk:"system_data"`
}
