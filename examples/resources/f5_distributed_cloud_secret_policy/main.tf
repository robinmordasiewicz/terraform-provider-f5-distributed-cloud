# Example configuration for f5_distributed_cloud_secret_policy

resource "f5_distributed_cloud_secret_policy" "example" {
  name        = "example-secret_policy"
  namespace   = "system"
  description = "Example SecretPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
