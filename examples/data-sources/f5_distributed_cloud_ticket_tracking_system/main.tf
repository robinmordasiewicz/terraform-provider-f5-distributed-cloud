# Example configuration for f5_distributed_cloud_ticket_tracking_system data source

data "f5_distributed_cloud_ticket_tracking_system" "example" {
  name      = "existing-ticket_tracking_system"
  namespace = "system"
}

output "ticket_tracking_system_id" {
  value = data.f5_distributed_cloud_ticket_tracking_system.example.id
}
