# Example configuration for f5xc_child_tenant data source

data "f5xc_child_tenant" "example" {
  name      = "existing-child_tenant"
  namespace = "system"
}

output "child_tenant_id" {
  value = data.f5xc_child_tenant.example.id
}
