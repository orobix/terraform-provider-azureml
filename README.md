# Terraform Provider for Azure Machine Learning

[![Actions Status](https://github.com/orobix/terraform-provider-azureml/workflows/Tests/badge.svg)](https://github.com/orobix/terraform-provider-azureml/actions)

Terraform provider for configuring [Azure Machine Learning](https://docs.microsoft.com/en-us/azure/machine-learning/)
Workspaces.

## Disclaimer

This Terraform provider relies on [Azure ML REST APIs](https://docs.microsoft.com/en-us/rest/api/azureml/)
which at the time of writing (26-02-2022) are still in preview and could thus introduce breaking changes from one
release to another. As such, it is possible that some features of the provider will stop working properly subsequently 
to Azure updates.

It also likely that once that a stable version of Azure ML APIs will be released the features offered by this provider
will be included in the official [azurerm provider](https://github.com/hashicorp/terraform-provider-azurerm) developed
by Hashicorp. Once that happens, this repository will not be maintained anymore.

## Usage examples

### Initialize the provider

```hcl
terraform {
  required_providers {
    azureml = {
      source  = "registry.terraform.io/orobix/azureml"
      version = "0.0.3"
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