# Example configuration for f5_distributed_cloud_waf_exclusion_policy

resource "f5_distributed_cloud_waf_exclusion_policy" "example" {
  name        = "example-waf_exclusion_policy"
  namespace   = "system"
  description = "Example WAFExclusionPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
