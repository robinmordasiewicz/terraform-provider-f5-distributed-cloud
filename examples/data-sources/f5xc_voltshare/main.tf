# Example configuration for f5xc_voltshare data source

data "f5xc_voltshare" "example" {
  name      = "existing-voltshare"
  namespace = "system"
}

output "voltshare_id" {
  value = data.f5xc_voltshare.example.id
}
