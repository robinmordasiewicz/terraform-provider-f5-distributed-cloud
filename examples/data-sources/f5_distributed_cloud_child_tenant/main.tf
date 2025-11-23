# Example configuration for f5_distributed_cloud_child_tenant data source

data "f5_distributed_cloud_child_tenant" "example" {
  name      = "existing-child_tenant"
  namespace = "system"
}

output "child_tenant_id" {
  value = data.f5_distributed_cloud_child_tenant.example.id
}
