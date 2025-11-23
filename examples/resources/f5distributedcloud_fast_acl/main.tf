# Example configuration for f5distributedcloud_fast_acl

resource "f5distributedcloud_fast_acl" "example" {
  name        = "example-fast_acl"
  namespace   = "system"
  description = "Example FastACL resource managed by Terraform"

  # Add additional configuration as needed
}
