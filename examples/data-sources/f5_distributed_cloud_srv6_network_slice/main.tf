# Example configuration for f5_distributed_cloud_srv6_network_slice data source

data "f5_distributed_cloud_srv6_network_slice" "example" {
  name      = "existing-srv6_network_slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.f5_distributed_cloud_srv6_network_slice.example.id
}
