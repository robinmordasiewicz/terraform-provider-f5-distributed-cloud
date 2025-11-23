# Example configuration for f5_distributed_cloud_ike2

resource "f5_distributed_cloud_ike2" "example" {
  name        = "example-ike2"
  namespace   = "system"
  description = "Example Ike2 resource managed by Terraform"

  # Add additional configuration as needed
}
