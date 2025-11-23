# Example configuration for f5distributedcloud_service_policy_set data source

data "f5distributedcloud_service_policy_set" "example" {
  name      = "existing-service_policy_set"
  namespace = "system"
}

output "service_policy_set_id" {
  value = data.f5distributedcloud_service_policy_set.example.id
}
