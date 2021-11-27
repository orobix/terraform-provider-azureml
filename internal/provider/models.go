package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SystemData struct {
	CreationDate     types.String `tfsdk:"creation_date"`
	CreationUser     types.String `tfsdk:"creation_user"`
	CreationUserType types.String `tfsdk:"creation_user_type"`

	LastModifiedDate     types.String `tfsdk:"last_modified_user_type"`
	LastModifiedUser     types.String `tfsdk:"last_modified_user_type"`
	LastModifiedUserType types.String `tfsdk:"last_modified_user_type"`
}

type DatastoreAuth struct {
	CredentialsType types.String `tfsdk:"credentials_type"`
	ClientId        types.String `tfsdk:"client_id"`
	ClientSecret    types.String `tfsdk:"client_secret"`
	AccountKey      types.String `tfsdk:"account_key"`
	SqlUserName     types.String `tfsdk:"sql_user_name"`
	SqlUserPassword types.String `tfsdk:"sql_user_password"`
}

type Datastore struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	StorageType          types.String `tfsdk:"storage_type"`
	StorageAccountName   types.String `tfsdk:"storage_account_name"`
	StorageContainerName types.String `tfsdk:"storage_container_name"`

	Auth       DatastoreAuth `tfsdk:"storage_container_name"`
	SystemData SystemData    `tfsdk:"storage_container_name"`
}
