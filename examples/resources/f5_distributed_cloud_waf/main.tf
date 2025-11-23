# Example: WAF
resource "f5_distributed_cloud_waf" "example" {
  name        = "my-waf"
  namespace   = "my-namespace"
  description = "Example Web Application Firewall configuration"
}
