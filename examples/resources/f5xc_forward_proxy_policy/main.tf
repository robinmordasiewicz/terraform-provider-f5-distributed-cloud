# Example configuration for f5xc_forward_proxy_policy

resource "f5xc_forward_proxy_policy" "example" {
  name        = "example-forward_proxy_policy"
  namespace   = "system"
  description = "Example ForwardProxyPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
