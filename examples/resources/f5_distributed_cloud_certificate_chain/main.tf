# Example configuration for f5_distributed_cloud_certificate_chain

resource "f5_distributed_cloud_certificate_chain" "example" {
  name        = "example-certificate_chain"
  namespace   = "system"
  description = "Example CertificateChain resource managed by Terraform"

  # Add additional configuration as needed
}
