# Example: Look up existing Active Service Policies
data "f5distributedcloud_active_service_policies" "example" {
  name      = "my-active-policies"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5distributedcloud_active_service_policies.example.enabled
}
