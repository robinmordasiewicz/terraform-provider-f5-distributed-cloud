# Example configuration for f5_distributed_cloud_allowed_tenant data source

data "f5_distributed_cloud_allowed_tenant" "example" {
  name      = "existing-allowed_tenant"
  namespace = "system"
}

output "allowed_tenant_id" {
  value = data.f5_distributed_cloud_allowed_tenant.example.id
}
