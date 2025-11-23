# Example configuration for f5distributedcloud_certificate_chain

resource "f5distributedcloud_certificate_chain" "example" {
  name        = "example-certificate_chain"
  namespace   = "system"
  description = "Example CertificateChain resource managed by Terraform"

  # Add additional configuration as needed
}
