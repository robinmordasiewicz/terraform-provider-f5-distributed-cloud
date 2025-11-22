# Example: Look up an existing secret policy
data "f5xc_secret_policy" "example" {
  name      = "my-secret-policy"
  namespace = "my-namespace"
}

# Output the enabled status
output "enabled" {
  value = data.f5xc_secret_policy.example.enabled
}

output "secret_policy_id" {
  value = data.f5xc_secret_policy.example.id
}
