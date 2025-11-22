# Example configuration for f5xc_route

resource "f5xc_route" "example" {
  name        = "example-route"
  namespace   = "system"
  description = "Example Route resource managed by Terraform"

  # Add additional configuration as needed
}
