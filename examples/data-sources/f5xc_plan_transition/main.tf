# Example configuration for f5xc_plan_transition data source

data "f5xc_plan_transition" "example" {
  name      = "existing-plan_transition"
  namespace = "system"
}

output "plan_transition_id" {
  value = data.f5xc_plan_transition.example.id
}
