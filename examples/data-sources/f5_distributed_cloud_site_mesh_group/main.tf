# Example configuration for f5_distributed_cloud_site_mesh_group data source

data "f5_distributed_cloud_site_mesh_group" "example" {
  name      = "existing-site_mesh_group"
  namespace = "system"
}

output "site_mesh_group_id" {
  value = data.f5_distributed_cloud_site_mesh_group.example.id
}
