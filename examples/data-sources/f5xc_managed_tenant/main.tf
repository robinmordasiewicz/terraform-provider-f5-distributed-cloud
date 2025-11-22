# Example configuration for f5xc_managed_tenant data source

data "f5xc_managed_tenant" "example" {
  name      = "existing-managed_tenant"
  namespace = "system"
}

output "managed_tenant_id" {
  value = data.f5xc_managed_tenant.example.id
}
