# Example configuration for f5_distributed_cloud_alert_template data source

data "f5_distributed_cloud_alert_template" "example" {
  name      = "existing-alert_template"
  namespace = "system"
}

output "alert_template_id" {
  value = data.f5_distributed_cloud_alert_template.example.id
}
