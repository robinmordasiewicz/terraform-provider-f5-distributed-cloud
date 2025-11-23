# Example configuration for f5_distributed_cloud_quota

resource "f5_distributed_cloud_quota" "example" {
  name        = "example-quota"
  namespace   = "system"
  description = "Example Quota resource managed by Terraform"

  # Add additional configuration as needed
}
