# Example configuration for f5xc_secret_policy data source

data "f5xc_secret_policy" "example" {
  name      = "existing-secret_policy"
  namespace = "system"
}

output "secret_policy_id" {
  value = data.f5xc_secret_policy.example.id
}
