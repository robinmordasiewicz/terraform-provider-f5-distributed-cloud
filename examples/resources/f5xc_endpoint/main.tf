# Example configuration for f5xc_endpoint

resource "f5xc_endpoint" "example" {
  name        = "example-endpoint"
  namespace   = "system"
  description = "Example Endpoint resource managed by Terraform"

  # Add additional configuration as needed
}
