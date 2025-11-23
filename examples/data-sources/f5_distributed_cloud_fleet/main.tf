# Example configuration for f5_distributed_cloud_fleet data source

data "f5_distributed_cloud_fleet" "example" {
  name      = "existing-fleet"
  namespace = "system"
}

output "fleet_id" {
  value = data.f5_distributed_cloud_fleet.example.id
}
