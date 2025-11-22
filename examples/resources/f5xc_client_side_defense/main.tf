# Example configuration for f5xc_client_side_defense

resource "f5xc_client_side_defense" "example" {
  name        = "example-client_side_defense"
  namespace   = "system"
  description = "Example ClientSideDefense resource managed by Terraform"

  # Add additional configuration as needed
}
