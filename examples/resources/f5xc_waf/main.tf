# Example: WAF
resource "f5xc_waf" "example" {
  name        = "my-waf"
  namespace   = "my-namespace"
  description = "Example Web Application Firewall configuration"
}
