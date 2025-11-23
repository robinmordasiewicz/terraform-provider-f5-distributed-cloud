# Example configuration for f5_distributed_cloud_tenant_profile data source

data "f5_distributed_cloud_tenant_profile" "example" {
  name      = "existing-tenant_profile"
  namespace = "system"
}

output "tenant_profile_id" {
  value = data.f5_distributed_cloud_tenant_profile.example.id
}
