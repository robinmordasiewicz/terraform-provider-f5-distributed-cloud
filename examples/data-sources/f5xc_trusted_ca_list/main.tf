# Example configuration for f5xc_trusted_ca_list data source

data "f5xc_trusted_ca_list" "example" {
  name      = "existing-trusted_ca_list"
  namespace = "system"
}

output "trusted_ca_list_id" {
  value = data.f5xc_trusted_ca_list.example.id
}
