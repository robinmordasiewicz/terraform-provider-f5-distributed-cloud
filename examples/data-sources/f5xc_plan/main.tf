# Example configuration for f5xc_plan data source

data "f5xc_plan" "example" {
  name      = "existing-plan"
  namespace = "system"
}

output "plan_id" {
  value = data.f5xc_plan.example.id
}
