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

resource "azureml_datastore" "test" {
  resource_group_name = "foo"
  workspace_name      = null
  name                = "test"
  description         = ""
  storage_type        = "AzureBlob"
  auth                = {
    credentials_type = ""
  }
}
