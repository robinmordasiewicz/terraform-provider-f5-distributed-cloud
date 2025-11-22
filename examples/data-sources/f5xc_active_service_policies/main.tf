# Example: Look up existing Active Service Policies
data "f5xc_active_service_policies" "example" {
  name      = "my-active-policies"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5xc_active_service_policies.example.enabled
}
