# Example configuration for f5xc_ticket_tracking_system

resource "f5xc_ticket_tracking_system" "example" {
  name        = "example-ticket_tracking_system"
  namespace   = "system"
  description = "Example TicketTrackingSystem resource managed by Terraform"

  # Add additional configuration as needed
}
