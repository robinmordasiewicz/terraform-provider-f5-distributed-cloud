# Example configuration for f5_distributed_cloud_managed_tenant data source

data "f5_distributed_cloud_managed_tenant" "example" {
  name      = "existing-managed_tenant"
  namespace = "system"
}

output "managed_tenant_id" {
  value = data.f5_distributed_cloud_managed_tenant.example.id
}
