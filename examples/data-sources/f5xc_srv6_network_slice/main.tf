# Example configuration for f5xc_srv6_network_slice data source

data "f5xc_srv6_network_slice" "example" {
  name      = "existing-srv6_network_slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.f5xc_srv6_network_slice.example.id
}
