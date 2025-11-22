# Example configuration for f5xc_public_ip data source

data "f5xc_public_ip" "example" {
  name      = "existing-public_ip"
  namespace = "system"
}

output "public_ip_id" {
  value = data.f5xc_public_ip.example.id
}
