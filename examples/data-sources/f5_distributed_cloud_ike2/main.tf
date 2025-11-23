# Example configuration for f5_distributed_cloud_ike2 data source

data "f5_distributed_cloud_ike2" "example" {
  name      = "existing-ike2"
  namespace = "system"
}

output "ike2_id" {
  value = data.f5_distributed_cloud_ike2.example.id
}
