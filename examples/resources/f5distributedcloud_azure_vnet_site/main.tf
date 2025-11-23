# Example configuration for f5distributedcloud_azure_vnet_site

resource "f5distributedcloud_azure_vnet_site" "example" {
  name        = "example-azure_vnet_site"
  namespace   = "system"
  description = "Example AzureVNETSite resource managed by Terraform"

  # Add additional configuration as needed
}
