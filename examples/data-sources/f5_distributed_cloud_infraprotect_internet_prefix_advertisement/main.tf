# Example configuration for f5_distributed_cloud_infraprotect_internet_prefix_advertisement data source

data "f5_distributed_cloud_infraprotect_internet_prefix_advertisement" "example" {
  name      = "existing-infraprotect_internet_prefix_advertisement"
  namespace = "system"
}

output "infraprotect_internet_prefix_advertisement_id" {
  value = data.f5_distributed_cloud_infraprotect_internet_prefix_advertisement.example.id
}
