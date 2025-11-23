# Example configuration for f5_distributed_cloud_contact

resource "f5_distributed_cloud_contact" "example" {
  name        = "example-contact"
  namespace   = "system"
  description = "Example Contact resource managed by Terraform"

  # Add additional configuration as needed
}
