# Example configuration for f5xc_irule

resource "f5xc_irule" "example" {
  name        = "example-irule"
  namespace   = "system"
  description = "Example Irule resource managed by Terraform"

  # Add additional configuration as needed
}
