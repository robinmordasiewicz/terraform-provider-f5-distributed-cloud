# Example configuration for f5xc_rrset

resource "f5xc_rrset" "example" {
  name        = "example-rrset"
  namespace   = "system"
  description = "Example Rrset resource managed by Terraform"

  # Add additional configuration as needed
}
