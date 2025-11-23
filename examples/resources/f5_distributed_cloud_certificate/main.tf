# Example configuration for f5_distributed_cloud_certificate

resource "f5_distributed_cloud_certificate" "example" {
  name        = "example-certificate"
  namespace   = "system"
  description = "Example Certificate resource managed by Terraform"

  # Add additional configuration as needed
}
