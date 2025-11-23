# Example configuration for f5_distributed_cloud_filter_set

resource "f5_distributed_cloud_filter_set" "example" {
  name        = "example-filter_set"
  namespace   = "system"
  description = "Example FilterSet resource managed by Terraform"

  # Add additional configuration as needed
}
