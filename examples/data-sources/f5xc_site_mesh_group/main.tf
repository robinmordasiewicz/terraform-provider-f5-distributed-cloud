# Example configuration for f5xc_site_mesh_group data source

data "f5xc_site_mesh_group" "example" {
  name      = "existing-site_mesh_group"
  namespace = "system"
}

output "site_mesh_group_id" {
  value = data.f5xc_site_mesh_group.example.id
}
