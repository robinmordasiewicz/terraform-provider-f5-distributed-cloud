# Example configuration for f5distributedcloud_public_ip data source

data "f5distributedcloud_public_ip" "example" {
  name      = "existing-public_ip"
  namespace = "system"
}

output "public_ip_id" {
  value = data.f5distributedcloud_public_ip.example.id
}
