# Example configuration for f5_distributed_cloud_voltshare data source

data "f5_distributed_cloud_voltshare" "example" {
  name      = "existing-voltshare"
  namespace = "system"
}

output "voltshare_id" {
  value = data.f5_distributed_cloud_voltshare.example.id
}
