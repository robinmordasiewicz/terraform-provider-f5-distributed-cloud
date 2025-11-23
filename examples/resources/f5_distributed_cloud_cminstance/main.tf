# Example configuration for f5_distributed_cloud_cminstance

resource "f5_distributed_cloud_cminstance" "example" {
  name        = "example-cminstance"
  namespace   = "system"
  description = "Example Cminstance resource managed by Terraform"

  # Add additional configuration as needed
}
