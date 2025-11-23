# Example configuration for f5_distributed_cloud_certified_hardware data source

data "f5_distributed_cloud_certified_hardware" "example" {
  name      = "existing-certified_hardware"
  namespace = "system"
}

output "certified_hardware_id" {
  value = data.f5_distributed_cloud_certified_hardware.example.id
}
