# Example configuration for f5xc_site_mesh_group

resource "f5xc_site_mesh_group" "example" {
  name        = "example-site_mesh_group"
  namespace   = "system"
  description = "Example SiteMeshGroup resource managed by Terraform"

  # Add additional configuration as needed
}
