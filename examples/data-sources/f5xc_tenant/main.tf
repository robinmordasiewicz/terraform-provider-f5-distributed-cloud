# Example configuration for f5xc_tenant data source

data "f5xc_tenant" "example" {
  name      = "existing-tenant"
  namespace = "system"
}

output "tenant_id" {
  value = data.f5xc_tenant.example.id
}
