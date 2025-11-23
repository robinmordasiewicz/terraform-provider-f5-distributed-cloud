# Example configuration for f5_distributed_cloud_site_mesh_group

resource "f5_distributed_cloud_site_mesh_group" "example" {
  name        = "example-site_mesh_group"
  namespace   = "system"
  description = "Example SiteMeshGroup resource managed by Terraform"

  # Add additional configuration as needed
}
