# Example configuration for f5xc_alert data source

data "f5xc_alert" "example" {
  name      = "existing-alert"
  namespace = "system"
}

output "alert_id" {
  value = data.f5xc_alert.example.id
}
