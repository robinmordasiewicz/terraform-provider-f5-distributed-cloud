# Example configuration for f5_distributed_cloud_ike1 data source

data "f5_distributed_cloud_ike1" "example" {
  name      = "existing-ike1"
  namespace = "system"
}

output "ike1_id" {
  value = data.f5_distributed_cloud_ike1.example.id
}
