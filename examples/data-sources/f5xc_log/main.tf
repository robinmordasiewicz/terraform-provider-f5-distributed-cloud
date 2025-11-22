# Example configuration for f5xc_log data source

data "f5xc_log" "example" {
  name      = "existing-log"
  namespace = "system"
}

output "log_id" {
  value = data.f5xc_log.example.id
}
