# Example configuration for f5_distributed_cloud_forwarding_class

resource "f5_distributed_cloud_forwarding_class" "example" {
  name        = "example-forwarding_class"
  namespace   = "system"
  description = "Example ForwardingClass resource managed by Terraform"

  # Add additional configuration as needed
}
