# Example configuration for f5xc_safe data source

data "f5xc_safe" "example" {
  name      = "existing-safe"
  namespace = "system"
}

output "safe_id" {
  value = data.f5xc_safe.example.id
}
