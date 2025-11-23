# Example configuration for f5distributedcloud_origin_pool

resource "f5distributedcloud_origin_pool" "example" {
  name        = "example-origin_pool"
  namespace   = "system"
  description = "Example OriginPool resource managed by Terraform"

  # Add additional configuration as needed
}
