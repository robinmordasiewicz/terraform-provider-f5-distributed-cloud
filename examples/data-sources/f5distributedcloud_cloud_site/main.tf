# Example: Look up an existing cloud site
data "f5distributedcloud_cloud_site" "example" {
  name      = "my-cloud-site"
  namespace = "system"
}

# Output the site state
output "site_state" {
  value = data.f5distributedcloud_cloud_site.example.site_state
}

output "site_id" {
  value = data.f5distributedcloud_cloud_site.example.id
}
