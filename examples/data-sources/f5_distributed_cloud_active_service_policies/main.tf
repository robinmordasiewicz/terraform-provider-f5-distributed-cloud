# Example: Look up existing Active Service Policies
data "f5_distributed_cloud_active_service_policies" "example" {
  name      = "my-active-policies"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5_distributed_cloud_active_service_policies.example.enabled
}
