# Example: Trusted CA List
resource "f5xc_trusted_ca_list" "example" {
  name        = "my-trusted-ca-list"
  namespace   = "my-namespace"
  description = "Example trusted CA list for certificate validation"
}
