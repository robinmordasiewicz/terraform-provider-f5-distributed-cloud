# Example: Look up an existing malicious user mitigation
data "f5xc_malicious_user_mitigation" "example" {
  name      = "my-mitigation"
  namespace = "my-namespace"
}

# Output the enabled status
output "enabled" {
  value = data.f5xc_malicious_user_mitigation.example.enabled
}

output "malicious_user_mitigation_id" {
  value = data.f5xc_malicious_user_mitigation.example.id
}
