# Example configuration for f5xc_advertise_policy data source

data "f5xc_advertise_policy" "example" {
  name      = "existing-advertise_policy"
  namespace = "system"
}

output "advertise_policy_id" {
  value = data.f5xc_advertise_policy.example.id
}
