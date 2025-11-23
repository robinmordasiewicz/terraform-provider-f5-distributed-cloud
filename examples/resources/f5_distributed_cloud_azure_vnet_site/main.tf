# Example configuration for f5_distributed_cloud_azure_vnet_site

resource "f5_distributed_cloud_azure_vnet_site" "example" {
  name        = "example-azure_vnet_site"
  namespace   = "system"
  description = "Example AzureVNETSite resource managed by Terraform"

  # Add additional configuration as needed
}
