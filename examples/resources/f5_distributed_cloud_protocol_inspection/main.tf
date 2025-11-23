# Example configuration for f5_distributed_cloud_protocol_inspection

resource "f5_distributed_cloud_protocol_inspection" "example" {
  name        = "example-protocol_inspection"
  namespace   = "system"
  description = "Example ProtocolInspection resource managed by Terraform"

  # Add additional configuration as needed
}
