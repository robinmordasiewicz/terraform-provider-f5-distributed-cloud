# Example configuration for f5distributedcloud_ip_prefix_set

resource "f5distributedcloud_ip_prefix_set" "example" {
  name        = "example-ip_prefix_set"
  namespace   = "system"
  description = "Example IPPrefixSet resource managed by Terraform"

  # Add additional configuration as needed
}
