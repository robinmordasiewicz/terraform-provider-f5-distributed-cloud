# Example configuration for f5xc_infraprotect_internet_prefix_advertisement data source

data "f5xc_infraprotect_internet_prefix_advertisement" "example" {
  name      = "existing-infraprotect_internet_prefix_advertisement"
  namespace = "system"
}

output "infraprotect_internet_prefix_advertisement_id" {
  value = data.f5xc_infraprotect_internet_prefix_advertisement.example.id
}
