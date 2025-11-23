# Example configuration for f5_distributed_cloud_user_identification

resource "f5_distributed_cloud_user_identification" "example" {
  name        = "example-user_identification"
  namespace   = "system"
  description = "Example UserIdentification resource managed by Terraform"

  # Add additional configuration as needed
}
