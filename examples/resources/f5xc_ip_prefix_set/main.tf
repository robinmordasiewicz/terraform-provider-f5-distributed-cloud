# Example configuration for f5xc_ip_prefix_set

resource "f5xc_ip_prefix_set" "example" {
  name        = "example-ip_prefix_set"
  namespace   = "system"
  description = "Example IPPrefixSet resource managed by Terraform"

  # Add additional configuration as needed
}
