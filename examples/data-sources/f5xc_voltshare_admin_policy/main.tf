# Example configuration for f5xc_voltshare_admin_policy data source

data "f5xc_voltshare_admin_policy" "example" {
  name      = "existing-voltshare_admin_policy"
  namespace = "system"
}

output "voltshare_admin_policy_id" {
  value = data.f5xc_voltshare_admin_policy.example.id
}
