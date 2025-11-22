# Example configuration for f5xc_allowed_tenant data source

data "f5xc_allowed_tenant" "example" {
  name      = "existing-allowed_tenant"
  namespace = "system"
}

output "allowed_tenant_id" {
  value = data.f5xc_allowed_tenant.example.id
}
