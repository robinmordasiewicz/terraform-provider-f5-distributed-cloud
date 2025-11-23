# Example configuration for f5_distributed_cloud_ike1

resource "f5_distributed_cloud_ike1" "example" {
  name        = "example-ike1"
  namespace   = "system"
  description = "Example Ike1 resource managed by Terraform"

  # Add additional configuration as needed
}
