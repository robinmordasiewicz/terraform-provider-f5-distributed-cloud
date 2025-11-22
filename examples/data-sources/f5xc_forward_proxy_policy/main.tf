# Example: Look up an existing forward proxy policy
data "f5xc_forward_proxy_policy" "example" {
  name      = "my-forward-proxy-policy"
  namespace = "my-namespace"
}

# Output the enabled status
output "enabled" {
  value = data.f5xc_forward_proxy_policy.example.enabled
}

output "forward_proxy_policy_id" {
  value = data.f5xc_forward_proxy_policy.example.id
}
