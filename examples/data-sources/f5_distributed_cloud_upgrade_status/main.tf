# Example configuration for f5_distributed_cloud_upgrade_status data source

data "f5_distributed_cloud_upgrade_status" "example" {
  name      = "existing-upgrade_status"
  namespace = "system"
}

output "upgrade_status_id" {
  value = data.f5_distributed_cloud_upgrade_status.example.id
}
