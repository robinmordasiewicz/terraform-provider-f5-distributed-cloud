# Example: Look up an existing GCP VPC site
data "f5xc_gcp_vpc_site" "example" {
  name      = "my-gcp-site"
  namespace = "system"
}

# Output the site state and GCP region
output "site_state" {
  value = data.f5xc_gcp_vpc_site.example.site_state
}

output "gcp_region" {
  value = data.f5xc_gcp_vpc_site.example.gcp_region
}
