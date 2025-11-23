# Example configuration for f5distributedcloud_site_mesh_group

resource "f5distributedcloud_site_mesh_group" "example" {
  name        = "example-site_mesh_group"
  namespace   = "system"
  description = "Example SiteMeshGroup resource managed by Terraform"

  # Add additional configuration as needed
}
