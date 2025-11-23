# Example configuration for f5distributedcloud_service_policy data source

data "f5distributedcloud_service_policy" "example" {
  name      = "existing-service_policy"
  namespace = "system"
}

output "service_policy_id" {
  value = data.f5distributedcloud_service_policy.example.id
}
