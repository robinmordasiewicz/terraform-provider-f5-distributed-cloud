# Example configuration for f5xc_infraprotect_information data source

data "f5xc_infraprotect_information" "example" {
  name      = "existing-infraprotect_information"
  namespace = "system"
}

output "infraprotect_information_id" {
  value = data.f5xc_infraprotect_information.example.id
}
