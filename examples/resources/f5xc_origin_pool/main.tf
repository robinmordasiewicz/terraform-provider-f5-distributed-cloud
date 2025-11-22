# Example configuration for f5xc_origin_pool

resource "f5xc_origin_pool" "example" {
  name        = "example-origin_pool"
  namespace   = "system"
  description = "Example OriginPool resource managed by Terraform"

  # Add additional configuration as needed
}
