# Example configuration for f5_distributed_cloud_fast_acl data source

data "f5_distributed_cloud_fast_acl" "example" {
  name      = "existing-fast_acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.f5_distributed_cloud_fast_acl.example.id
}
