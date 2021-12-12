terraform {
  required_providers {
    azureml = {
      source = "registry.terraform.io/Telemaco019/azureml"
    }
  }
}

provider "azureml" {
  client_id       = ""
  client_secret   = ""
  tenant_id       = ""
  subscription_id = ""
}
