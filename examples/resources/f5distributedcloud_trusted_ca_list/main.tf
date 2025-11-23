# Example configuration for f5distributedcloud_trusted_ca_list

resource "f5distributedcloud_trusted_ca_list" "example" {
  name        = "example-trusted_ca_list"
  namespace   = "system"
  description = "Example TrustedCAList resource managed by Terraform"

  # Add additional configuration as needed
}
