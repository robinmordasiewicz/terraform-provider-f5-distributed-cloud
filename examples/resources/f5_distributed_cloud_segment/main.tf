# Example configuration for f5_distributed_cloud_segment

resource "f5_distributed_cloud_segment" "example" {
  name        = "example-segment"
  namespace   = "system"
  description = "Example Segment resource managed by Terraform"

  # Add additional configuration as needed
}
