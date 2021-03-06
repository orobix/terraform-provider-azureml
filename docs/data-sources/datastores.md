---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "azureml_datastores Data Source - terraform-provider-azureml"
subcategory: ""
description: |-
  Use this resource to retrieve the list of Datastores of a certain Azure ML Workspace. Authentication credentials are not included in the provided information.
---

# azureml_datastores (Data Source)

Use this resource to retrieve the list of Datastores of a certain Azure ML Workspace. Authentication credentials are not included in the provided information.

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

- **resource_group_name** (String) The name of the resource group of the Azure ML Workspace to which the datastore belongs to.
- **workspace_name** (String) The name of the Azure ML Workspace to which the datastore belongs to.

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **datastores** (List of Object) (see [below for nested schema](#nestedatt--datastores))

<a id="nestedatt--datastores"></a>
### Nested Schema for `datastores`

Read-Only:

- **creation_date** (String)
- **creation_user** (String)
- **creation_user_type** (String)
- **credentials_type** (String)
- **description** (String)
- **id** (String)
- **is_default** (Boolean)
- **last_modified_date** (String)
- **last_modified_user** (String)
- **last_modified_user_type** (String)
- **name** (String)
- **resource_group_name** (String)
- **storage_account_name** (String)
- **storage_container_name** (String)
- **storage_type** (String)
- **workspace_name** (String)


