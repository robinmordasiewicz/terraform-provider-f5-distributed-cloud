# Example configuration for f5_distributed_cloud_scim

resource "f5_distributed_cloud_scim" "example" {
  name        = "example-scim"
  namespace   = "system"
  description = "Example Scim resource managed by Terraform"

  # Add additional configuration as needed
}
