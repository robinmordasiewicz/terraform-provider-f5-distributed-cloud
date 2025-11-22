# Example configuration for f5xc_alert_receiver data source

data "f5xc_alert_receiver" "example" {
  name      = "existing-alert_receiver"
  namespace = "system"
}

output "alert_receiver_id" {
  value = data.f5xc_alert_receiver.example.id
}
