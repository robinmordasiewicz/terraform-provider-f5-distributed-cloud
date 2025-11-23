# Example configuration for f5_distributed_cloud_trusted_ca_list data source

data "f5_distributed_cloud_trusted_ca_list" "example" {
  name      = "existing-trusted_ca_list"
  namespace = "system"
}

output "trusted_ca_list_id" {
  value = data.f5_distributed_cloud_trusted_ca_list.example.id
}
