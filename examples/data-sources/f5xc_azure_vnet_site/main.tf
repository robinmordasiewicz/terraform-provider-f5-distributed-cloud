# Example: Look up an existing Azure VNET site
data "f5xc_azure_vnet_site" "example" {
  name      = "my-azure-site"
  namespace = "system"
}

# Output the site state and Azure region
output "site_state" {
  value = data.f5xc_azure_vnet_site.example.site_state
}

output "azure_region" {
  value = data.f5xc_azure_vnet_site.example.azure_region
}
