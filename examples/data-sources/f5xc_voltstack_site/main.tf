# Example: Look up an existing Voltstack site
data "f5xc_voltstack_site" "example" {
  name      = "my-voltstack-site"
  namespace = "system"
}

# Output the site state and type
output "site_state" {
  value = data.f5xc_voltstack_site.example.site_state
}

output "site_type" {
  value = data.f5xc_voltstack_site.example.site_type
}
