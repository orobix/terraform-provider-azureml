terraform {
  required_providers {
    azureml = {
      source = "registry.terraform.io/orobix/azureml"
    }
  }
}

provider "azureml" {
  client_id       = ""
  client_secret   = ""
  tenant_id       = ""
  subscription_id = ""
}
