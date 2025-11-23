# Example configuration for f5_distributed_cloud_report data source

data "f5_distributed_cloud_report" "example" {
  name      = "existing-report"
  namespace = "system"
}

output "report_id" {
  value = data.f5_distributed_cloud_report.example.id
}
