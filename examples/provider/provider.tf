terraform {
  required_providers {
    azureml = {
      source  = "registry.terraform.io/Telemaco019/azureml"
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

data "azureml_datastore" "datastore" {
  resource_group_name = var.resource_group_name
  workspace_name      = var.workspace_name
  name                = "test"
}
output "datastore" {
  value = data.azureml_datastore.datastore
}

data "azureml_datastores" "datastores" {
  resource_group_name = var.resource_group_name
  workspace_name      = var.workspace_name
}

resource "azureml_datastore" "example" {
  resource_group_name = var.resource_group_name
  workspace_name      = var.workspace_name
  name                = "example2"
  description         = "example"
  storage_type        = "AzureBlob"

  storage_account_name   = "pippo"
  storage_container_name = "example"

  auth {
    credentials_type = "ServicePrincipal"
    client_id        = var.client_id
    client_secret    = var.client_secret
    tenant_id        = var.tenant_id
  }
}
