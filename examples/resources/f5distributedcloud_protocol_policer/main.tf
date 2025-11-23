# Example configuration for f5distributedcloud_protocol_policer

resource "f5distributedcloud_protocol_policer" "example" {
  name        = "example-protocol_policer"
  namespace   = "system"
  description = "Example ProtocolPolicer resource managed by Terraform"

  # Add additional configuration as needed
}
