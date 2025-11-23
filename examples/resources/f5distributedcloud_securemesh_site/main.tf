# Example configuration for f5distributedcloud_securemesh_site

resource "f5distributedcloud_securemesh_site" "example" {
  name        = "example-securemesh_site"
  namespace   = "system"
  description = "Example SecuremeshSite resource managed by Terraform"

  # Add additional configuration as needed
}
