# Example configuration for f5xc_securemesh_site data source

data "f5xc_securemesh_site" "example" {
  name      = "existing-securemesh_site"
  namespace = "system"
}

output "securemesh_site_id" {
  value = data.f5xc_securemesh_site.example.id
}
