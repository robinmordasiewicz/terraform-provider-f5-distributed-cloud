# Example configuration for f5xc_infraprotect data source

data "f5xc_infraprotect" "example" {
  name      = "existing-infraprotect"
  namespace = "system"
}

output "infraprotect_id" {
  value = data.f5xc_infraprotect.example.id
}
