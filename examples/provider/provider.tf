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

output "test" {
  value = data.azureml_datastores.test
}