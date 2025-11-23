# Example configuration for f5_distributed_cloud_bigip_irule

resource "f5_distributed_cloud_bigip_irule" "example" {
  name        = "example-bigip_irule"
  namespace   = "system"
  description = "Example BigipIrule resource managed by Terraform"

  # Add additional configuration as needed
}
