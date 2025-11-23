# Example configuration for f5distributedcloud_workload data source

data "f5distributedcloud_workload" "example" {
  name      = "existing-workload"
  namespace = "system"
}

output "workload_id" {
  value = data.f5distributedcloud_workload.example.id
}
