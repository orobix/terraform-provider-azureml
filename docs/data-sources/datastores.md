---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "azureml_datastores Data Source - terraform-provider-azureml"
subcategory: ""
description: |-
  Use this resource to fetch the list of of the Datastores belonging to a certain Azure ML Workspace. Authentication credentials are not included in the returned datastore list.
---

# azureml_datastores (Data Source)

Use this resource to fetch the list of of the Datastores belonging to a certain Azure ML Workspace. Authentication credentials are not included in the returned datastore list.

## Example Usage

```terraform
data "azureml_datastore" "example" {
  resource_group_name = "example"
  workspace_name      = "example"
  name                = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **resource_group_name** (String) The name of the resource group of the Azure ML Workspace.
- **workspace_name** (String) The name of the Azure ML Workspace from which the datastores will be fetched.

### Read-Only

- **datastores** (Attributes Set) (see [below for nested schema](#nestedatt--datastores))

<a id="nestedatt--datastores"></a>
### Nested Schema for `datastores`

Read-Only:

- **credentials_type** (String) The type of credentials used for authenticating with the underlying storage. Possible values are: [AccountKey Certificate None Sas ServicePrincipal SqlAdmin].
- **description** (String) The description of the datastore.
- **id** (String) The ID of the datastore.
- **is_default** (Boolean) Is the datastore the default datastore of the Azure ML Workspace?
- **name** (String) The name of the datastore.
- **resource_group_name** (String) The name of the resource group of the Azure ML Workspace to which the datastore belongs to.
- **storage_account_name** (String) The name of the Storage Account to which the datastore is linked to.
- **storage_container_name** (String) The name of the Storage Container to which the datastore is linked to.
- **storage_type** (String) The type of the storage to which the datstore is linked to. Possible values are: [AzureFile AzureBlob AzureDataLakeGen1 AzureDataLakeGen2 AzureMySql AzurePostgreSql AzureSqlDatabase GlusterFs]
- **system_data** (Attributes) (see [below for nested schema](#nestedatt--datastores--system_data))
- **workspace_name** (String) The name of the Azure ML Workspace to which the datastore belongs to.

<a id="nestedatt--datastores--system_data"></a>
### Nested Schema for `datastores.system_data`

Read-Only:

- **creation_date** (String) The timestamp corresponding to the creation of the datastore.
- **creation_user** (String) The user that created the datastore.
- **creation_user_type** (String) The kind of user that created the datastore (Service Principal or User).
- **last_modified_date** (String) The timestamp corresponding to the last update of the datastore.
- **last_modified_user** (String) The user that last updated the datastore.
- **last_modified_user_type** (String) The kind of user that last updated the datastore (Service Principal or User).

