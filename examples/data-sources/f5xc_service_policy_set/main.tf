# Example configuration for f5xc_service_policy_set data source

data "f5xc_service_policy_set" "example" {
  name      = "existing-service_policy_set"
  namespace = "system"
}

output "service_policy_set_id" {
  value = data.f5xc_service_policy_set.example.id
}
