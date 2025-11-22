# Example configuration for f5xc_workload_flavor data source

data "f5xc_workload_flavor" "example" {
  name      = "existing-workload_flavor"
  namespace = "system"
}

output "workload_flavor_id" {
  value = data.f5xc_workload_flavor.example.id
}
