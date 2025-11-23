# Example configuration for f5distributedcloud_azure_vnet_site data source

data "f5distributedcloud_azure_vnet_site" "example" {
  name      = "existing-azure_vnet_site"
  namespace = "system"
}

output "azure_vnet_site_id" {
  value = data.f5distributedcloud_azure_vnet_site.example.id
}
