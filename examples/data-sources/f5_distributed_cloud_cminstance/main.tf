# Example configuration for f5_distributed_cloud_cminstance data source

data "f5_distributed_cloud_cminstance" "example" {
  name      = "existing-cminstance"
  namespace = "system"
}

output "cminstance_id" {
  value = data.f5_distributed_cloud_cminstance.example.id
}
