# Example: CRL (Certificate Revocation List)
resource "f5xc_crl" "example" {
  name        = "my-crl"
  namespace   = "my-namespace"
  description = "Example certificate revocation list"
}
