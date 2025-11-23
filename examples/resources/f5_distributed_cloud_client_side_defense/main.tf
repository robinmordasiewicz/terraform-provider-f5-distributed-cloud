# Example configuration for f5_distributed_cloud_client_side_defense

resource "f5_distributed_cloud_client_side_defense" "example" {
  name        = "example-client_side_defense"
  namespace   = "system"
  description = "Example ClientSideDefense resource managed by Terraform"

  # Add additional configuration as needed
}
