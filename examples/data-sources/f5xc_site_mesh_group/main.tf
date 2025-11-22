# Example: Look up an existing site mesh group
data "f5xc_site_mesh_group" "example" {
  name      = "my-site-mesh-group"
  namespace = "system"
}

# Output the mesh type
output "mesh_type" {
  value = data.f5xc_site_mesh_group.example.mesh_type
}

output "site_mesh_group_id" {
  value = data.f5xc_site_mesh_group.example.id
}
