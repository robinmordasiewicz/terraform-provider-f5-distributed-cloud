# Example: Certificate Chain
resource "f5xc_certificate_chain" "example" {
  name        = "my-certificate-chain"
  namespace   = "my-namespace"
  description = "Example certificate chain for TLS configuration"
}
