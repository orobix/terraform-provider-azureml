package provider

import (
	"github.com/Telemaco019/azureml-go-sdk/workspace"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mergeToDatastoreAuth(tfAuth *DatastoreAuth, clientAuth *workspace.DatastoreAuth) DatastoreAuth {
	auth := DatastoreAuth{
		CredentialsType: types.String{Value: clientAuth.CredentialsType},
		// read from resource data since APIs do not return secrets
		ClientSecret:    tfAuth.ClientSecret,
		AccountKey:      tfAuth.AccountKey,
		SqlUserPassword: tfAuth.SqlUserPassword,
	}
	if tfAuth.ClientId.Null == true {
		auth.ClientId = types.String{Null: true}
	} else {
		auth.ClientId = types.String{Value: clientAuth.ClientId}
	}
	if tfAuth.TenantId.Null == true {
		auth.TenantId = types.String{Null: true}
	} else {
		auth.TenantId = types.String{Value: clientAuth.TenantId}
	}
	if tfAuth.SqlUserName.Null == true {
		auth.SqlUserName = types.String{Null: true}
	} else {
		auth.SqlUserName = types.String{Value: clientAuth.SqlUserName}
	}
	return auth
}

func mergeToDatastoreWithAuth(tfDatastore *DatastoreWithAuth, clientDatastore *workspace.Datastore) DatastoreWithAuth {
	result := new(DatastoreWithAuth)

	result.ResourceGroupName = tfDatastore.ResourceGroupName
	result.WorkspaceName = tfDatastore.WorkspaceName

	result.ID = types.String{Value: clientDatastore.Id}
	result.Name = types.String{Value: clientDatastore.Name}
	result.Description = types.String{Value: clientDatastore.Description}
	result.IsDefault = types.Bool{Value: clientDatastore.IsDefault}
	result.StorageType = types.String{Value: clientDatastore.StorageType}
	result.StorageAccountName = types.String{Value: clientDatastore.StorageAccountName}
	result.StorageContainerName = types.String{Value: clientDatastore.StorageContainerName}
	result.Auth = mergeToDatastoreAuth(&tfDatastore.Auth, &clientDatastore.Auth)
	result.SystemData = SystemData{
		CreationDate:         types.String{Value: clientDatastore.SystemData.CreationDate.Format(defaultDateFormat)},
		CreationUser:         types.String{Value: clientDatastore.SystemData.CreationUser},
		CreationUserType:     types.String{Value: clientDatastore.SystemData.CreationUserType},
		LastModifiedDate:     types.String{Value: clientDatastore.SystemData.LastModifiedDate.Format(defaultDateFormat)},
		LastModifiedUser:     types.String{Value: clientDatastore.SystemData.LastModifiedUser},
		LastModifiedUserType: types.String{Value: clientDatastore.SystemData.LastModifiedUserType},
	}

	return *result
}
