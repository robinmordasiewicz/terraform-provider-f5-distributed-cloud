# Example configuration for f5xc_segment

resource "f5xc_segment" "example" {
  name        = "example-segment"
  namespace   = "system"
  description = "Example Segment resource managed by Terraform"

  # Add additional configuration as needed
}
