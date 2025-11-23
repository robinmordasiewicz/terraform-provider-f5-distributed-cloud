# Example configuration for f5_distributed_cloud_ip_prefix_set

resource "f5_distributed_cloud_ip_prefix_set" "example" {
  name        = "example-ip_prefix_set"
  namespace   = "system"
  description = "Example IPPrefixSet resource managed by Terraform"

  # Add additional configuration as needed
}
