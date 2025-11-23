# Example: WAF
resource "f5distributedcloud_waf" "example" {
  name        = "my-waf"
  namespace   = "my-namespace"
  description = "Example Web Application Firewall configuration"
}
