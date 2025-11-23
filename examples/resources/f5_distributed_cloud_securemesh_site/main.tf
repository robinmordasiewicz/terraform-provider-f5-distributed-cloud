# Example configuration for f5_distributed_cloud_securemesh_site

resource "f5_distributed_cloud_securemesh_site" "example" {
  name        = "example-securemesh_site"
  namespace   = "system"
  description = "Example SecuremeshSite resource managed by Terraform"

  # Add additional configuration as needed
}
