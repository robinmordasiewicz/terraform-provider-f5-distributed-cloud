# Example configuration for f5xc_service_policy data source

data "f5xc_service_policy" "example" {
  name      = "existing-service_policy"
  namespace = "system"
}

output "service_policy_id" {
  value = data.f5xc_service_policy.example.id
}
