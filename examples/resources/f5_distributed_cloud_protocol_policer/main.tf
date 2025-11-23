# Example configuration for f5_distributed_cloud_protocol_policer

resource "f5_distributed_cloud_protocol_policer" "example" {
  name        = "example-protocol_policer"
  namespace   = "system"
  description = "Example ProtocolPolicer resource managed by Terraform"

  # Add additional configuration as needed
}
