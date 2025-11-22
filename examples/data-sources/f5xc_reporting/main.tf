# Example configuration for f5xc_reporting data source

data "f5xc_reporting" "example" {
  name      = "existing-reporting"
  namespace = "system"
}

output "reporting_id" {
  value = data.f5xc_reporting.example.id
}
