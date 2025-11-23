# Example configuration for f5distributedcloud_forward_proxy_policy data source

data "f5distributedcloud_forward_proxy_policy" "example" {
  name      = "existing-forward_proxy_policy"
  namespace = "system"
}

output "forward_proxy_policy_id" {
  value = data.f5distributedcloud_forward_proxy_policy.example.id
}
