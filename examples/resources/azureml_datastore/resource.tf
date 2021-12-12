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
