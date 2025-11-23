# Example configuration for f5_distributed_cloud_plan data source

data "f5_distributed_cloud_plan" "example" {
  name      = "existing-plan"
  namespace = "system"
}

output "plan_id" {
  value = data.f5_distributed_cloud_plan.example.id
}
