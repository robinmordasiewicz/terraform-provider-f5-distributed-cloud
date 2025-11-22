# Example configuration for f5xc_malicious_user_mitigation data source

data "f5xc_malicious_user_mitigation" "example" {
  name      = "existing-malicious_user_mitigation"
  namespace = "system"
}

output "malicious_user_mitigation_id" {
  value = data.f5xc_malicious_user_mitigation.example.id
}
