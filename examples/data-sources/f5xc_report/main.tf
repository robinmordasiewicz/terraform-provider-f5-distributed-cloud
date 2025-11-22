# Example configuration for f5xc_report data source

data "f5xc_report" "example" {
  name      = "existing-report"
  namespace = "system"
}

output "report_id" {
  value = data.f5xc_report.example.id
}
