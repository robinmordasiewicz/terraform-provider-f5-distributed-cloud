# Example configuration for f5_distributed_cloud_bgp data source

data "f5_distributed_cloud_bgp" "example" {
  name      = "existing-bgp"
  namespace = "system"
}

output "bgp_id" {
  value = data.f5_distributed_cloud_bgp.example.id
}
