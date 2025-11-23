# Example configuration for f5_distributed_cloud_registration

resource "f5_distributed_cloud_registration" "example" {
  name        = "example-registration"
  namespace   = "system"
  description = "Example Registration resource managed by Terraform"

  # Add additional configuration as needed
}
