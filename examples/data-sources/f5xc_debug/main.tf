# Example configuration for f5xc_debug data source

data "f5xc_debug" "example" {
  name      = "existing-debug"
  namespace = "system"
}

output "debug_id" {
  value = data.f5xc_debug.example.id
}
