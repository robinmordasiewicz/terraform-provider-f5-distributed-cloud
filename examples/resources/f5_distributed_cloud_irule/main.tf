# Example configuration for f5_distributed_cloud_irule

resource "f5_distributed_cloud_irule" "example" {
  name        = "example-irule"
  namespace   = "system"
  description = "Example Irule resource managed by Terraform"

  # Add additional configuration as needed
}
