# Example configuration for f5_distributed_cloud_reporting data source

data "f5_distributed_cloud_reporting" "example" {
  name      = "existing-reporting"
  namespace = "system"
}

output "reporting_id" {
  value = data.f5_distributed_cloud_reporting.example.id
}
