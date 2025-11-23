# Example configuration for f5_distributed_cloud_scim data source

data "f5_distributed_cloud_scim" "example" {
  name      = "existing-scim"
  namespace = "system"
}

output "scim_id" {
  value = data.f5_distributed_cloud_scim.example.id
}
