# Example configuration for f5_distributed_cloud_tenant data source

data "f5_distributed_cloud_tenant" "example" {
  name      = "existing-tenant"
  namespace = "system"
}

output "tenant_id" {
  value = data.f5_distributed_cloud_tenant.example.id
}
