# Example configuration for f5_distributed_cloud_plan_transition data source

data "f5_distributed_cloud_plan_transition" "example" {
  name      = "existing-plan_transition"
  namespace = "system"
}

output "plan_transition_id" {
  value = data.f5_distributed_cloud_plan_transition.example.id
}
