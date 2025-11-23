# Example configuration for f5distributedcloud_segment

resource "f5distributedcloud_segment" "example" {
  name        = "example-segment"
  namespace   = "system"
  description = "Example Segment resource managed by Terraform"

  # Add additional configuration as needed
}
