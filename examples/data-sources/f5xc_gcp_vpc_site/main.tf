# Example configuration for f5xc_gcp_vpc_site data source

data "f5xc_gcp_vpc_site" "example" {
  name      = "existing-gcp_vpc_site"
  namespace = "system"
}

output "gcp_vpc_site_id" {
  value = data.f5xc_gcp_vpc_site.example.id
}
