# Example configuration for f5_distributed_cloud_crl

resource "f5_distributed_cloud_crl" "example" {
  name        = "example-crl"
  namespace   = "system"
  description = "Example CRL resource managed by Terraform"

  # Add additional configuration as needed
}
