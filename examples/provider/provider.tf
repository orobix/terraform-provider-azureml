terraform {
  required_providers {
    azureml = {
      source  = "registry.terraform.io/zanotti/azureml"
      version = "0.0.1"
    }
  }
}

provider "azureml" {
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
  subscription_id = var.subscription_id
}

data "azureml_datastores" "test" {
  resource_group_name = var.resource_group_name
  workspace_name      = var.workspace_name
}

data "azureml_datastore" "test" {
  resource_group_name = var.resource_group_name
  workspace_name      = var.workspace_name
  name                = "workspaceworkingdirectory"
}

output "all_datastores" {
  value = data.azureml_datastores.test
}
output "single_datastore" {
  value = data.azureml_datastore.test
}


resource "azureml_datastore" "example" {
  resource_group_name = "example"
  workspace_name      = "example"
  name                = "example"
  description         = "example"
  storage_type        = "AzureBlob"

  storage_account_name   = "example"
  storage_container_name = "example"

  auth = {
    credentials_type = "AzureDataLakeGen2"
    client_id        = "client-id"
    client_secret    = "client-secret"
    tenant_id        = "tenant-id"
  }
}
