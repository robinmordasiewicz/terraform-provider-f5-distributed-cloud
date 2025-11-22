# Example configuration for f5xc_ticket_tracking_system data source

data "f5xc_ticket_tracking_system" "example" {
  name      = "existing-ticket_tracking_system"
  namespace = "system"
}

output "ticket_tracking_system_id" {
  value = data.f5xc_ticket_tracking_system.example.id
}
