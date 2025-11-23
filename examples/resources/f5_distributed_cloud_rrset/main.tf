# Example configuration for f5_distributed_cloud_rrset

resource "f5_distributed_cloud_rrset" "example" {
  name        = "example-rrset"
  namespace   = "system"
  description = "Example Rrset resource managed by Terraform"

  # Add additional configuration as needed
}
