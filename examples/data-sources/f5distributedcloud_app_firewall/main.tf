# Example configuration for f5distributedcloud_app_firewall data source

data "f5distributedcloud_app_firewall" "example" {
  name      = "existing-app_firewall"
  namespace = "system"
}

output "app_firewall_id" {
  value = data.f5distributedcloud_app_firewall.example.id
}
