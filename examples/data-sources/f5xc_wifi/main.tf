# Example configuration for f5xc_wifi data source

data "f5xc_wifi" "example" {
  name      = "existing-wifi"
  namespace = "system"
}

output "wifi_id" {
  value = data.f5xc_wifi.example.id
}
