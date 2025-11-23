# Example configuration for f5_distributed_cloud_infraprotect data source

data "f5_distributed_cloud_infraprotect" "example" {
  name      = "existing-infraprotect"
  namespace = "system"
}

output "infraprotect_id" {
  value = data.f5_distributed_cloud_infraprotect.example.id
}
