# Example configuration for f5_distributed_cloud_apm data source

data "f5_distributed_cloud_apm" "example" {
  name      = "existing-apm"
  namespace = "system"
}

output "apm_id" {
  value = data.f5_distributed_cloud_apm.example.id
}
