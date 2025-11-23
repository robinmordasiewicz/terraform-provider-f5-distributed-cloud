# Example configuration for f5_distributed_cloud_receiver

resource "f5_distributed_cloud_receiver" "example" {
  name        = "example-receiver"
  namespace   = "system"
  description = "Example Receiver resource managed by Terraform"

  # Add additional configuration as needed
}
