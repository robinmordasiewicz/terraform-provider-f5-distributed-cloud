# Example configuration for f5xc_quota data source

data "f5xc_quota" "example" {
  name      = "existing-quota"
  namespace = "system"
}

output "quota_id" {
  value = data.f5xc_quota.example.id
}
