# Example configuration for f5xc_alert_template data source

data "f5xc_alert_template" "example" {
  name      = "existing-alert_template"
  namespace = "system"
}

output "alert_template_id" {
  value = data.f5xc_alert_template.example.id
}
