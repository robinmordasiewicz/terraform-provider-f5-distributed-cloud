# Example configuration for f5xc_rbac_policy data source

data "f5xc_rbac_policy" "example" {
  name      = "existing-rbac_policy"
  namespace = "system"
}

output "rbac_policy_id" {
  value = data.f5xc_rbac_policy.example.id
}
