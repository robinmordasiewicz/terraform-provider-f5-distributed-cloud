# Example configuration for f5xc_virtual_host data source

data "f5xc_virtual_host" "example" {
  name      = "existing-virtual_host"
  namespace = "system"
}

output "virtual_host_id" {
  value = data.f5xc_virtual_host.example.id
}
