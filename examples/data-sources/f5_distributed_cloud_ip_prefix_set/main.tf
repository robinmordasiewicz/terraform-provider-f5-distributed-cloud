# Example configuration for f5_distributed_cloud_ip_prefix_set data source

data "f5_distributed_cloud_ip_prefix_set" "example" {
  name      = "existing-ip_prefix_set"
  namespace = "system"
}

output "ip_prefix_set_id" {
  value = data.f5_distributed_cloud_ip_prefix_set.example.id
}
