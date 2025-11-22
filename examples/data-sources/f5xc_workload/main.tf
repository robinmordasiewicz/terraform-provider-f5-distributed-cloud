# Example configuration for f5xc_workload data source

data "f5xc_workload" "example" {
  name      = "existing-workload"
  namespace = "system"
}

output "workload_id" {
  value = data.f5xc_workload.example.id
}
