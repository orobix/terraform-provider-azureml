# Terraform Provider for Azure Machine Learning

Terraform provider for managing [Azure Machine Learning](https://docs.microsoft.com/en-us/azure/machine-learning/) Workspaces. 

## Usage examples

### Initialize the provider
```hcl
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
```

### Configure a datastore
```hcl
resource "azureml_datastore" "example" {
  resource_group_name = "rg-name"
  workspace_name      = "ws-name"
  name                = "example"
  description         = "example"
  storage_type        = "AzureBlob"

  storage_account_name   = "example"
  storage_container_name = "example"

  auth {
    credentials_type = "ServicePrincipal"
    client_id        = var.client_id
    client_secret    = var.client_secret
    tenant_id        = var.tenant_id
  }
}
```

## Provider development quickstart

### Build the provider
```shell
make build
```

### Install the provider on your local machine
```shell
make install
```

### Generate the provider documentation
```shell
go generate
```