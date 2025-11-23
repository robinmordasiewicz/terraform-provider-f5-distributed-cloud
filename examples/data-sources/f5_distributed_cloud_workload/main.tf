# Example configuration for f5_distributed_cloud_workload data source

data "f5_distributed_cloud_workload" "example" {
  name      = "existing-workload"
  namespace = "system"
}

output "workload_id" {
  value = data.f5_distributed_cloud_workload.example.id
}
