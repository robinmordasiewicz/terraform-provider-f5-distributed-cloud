# Example: Proxy
resource "f5xc_proxy" "example" {
  name        = "my-proxy"
  namespace   = "my-namespace"
  description = "Example proxy configuration"
}
