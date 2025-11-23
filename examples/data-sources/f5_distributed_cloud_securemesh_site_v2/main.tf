# Example configuration for f5_distributed_cloud_securemesh_site_v2 data source

data "f5_distributed_cloud_securemesh_site_v2" "example" {
  name      = "existing-securemesh_site_v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.f5_distributed_cloud_securemesh_site_v2.example.id
}
