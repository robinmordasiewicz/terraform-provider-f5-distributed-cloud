# Example configuration for f5xc_secret_policy

resource "f5xc_secret_policy" "example" {
  name        = "example-secret_policy"
  namespace   = "system"
  description = "Example SecretPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
