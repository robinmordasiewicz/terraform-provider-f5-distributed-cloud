# Example configuration for f5xc_fast_acl

resource "f5xc_fast_acl" "example" {
  name        = "example-fast_acl"
  namespace   = "system"
  description = "Example FastACL resource managed by Terraform"

  # Add additional configuration as needed
}
