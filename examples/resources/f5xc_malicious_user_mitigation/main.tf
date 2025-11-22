# Example configuration for f5xc_malicious_user_mitigation

resource "f5xc_malicious_user_mitigation" "example" {
  name        = "example-malicious_user_mitigation"
  namespace   = "system"
  description = "Example MaliciousUserMitigation resource managed by Terraform"

  # Add additional configuration as needed
}
