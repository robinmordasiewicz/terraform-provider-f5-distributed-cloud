# Example configuration for f5_distributed_cloud_service_policy

resource "f5_distributed_cloud_service_policy" "example" {
  name        = "example-service_policy"
  namespace   = "system"
  description = "Example ServicePolicy resource managed by Terraform"

  # Add additional configuration as needed
}
