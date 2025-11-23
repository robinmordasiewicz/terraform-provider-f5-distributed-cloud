# Example configuration for f5_distributed_cloud_securemesh_site data source

data "f5_distributed_cloud_securemesh_site" "example" {
  name      = "existing-securemesh_site"
  namespace = "system"
}

output "securemesh_site_id" {
  value = data.f5_distributed_cloud_securemesh_site.example.id
}
