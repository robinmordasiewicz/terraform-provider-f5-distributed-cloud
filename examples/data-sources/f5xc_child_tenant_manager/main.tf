# Example configuration for f5xc_child_tenant_manager data source

data "f5xc_child_tenant_manager" "example" {
  name      = "existing-child_tenant_manager"
  namespace = "system"
}

output "child_tenant_manager_id" {
  value = data.f5xc_child_tenant_manager.example.id
}
