# Example configuration for f5xc_bgp data source

data "f5xc_bgp" "example" {
  name      = "existing-bgp"
  namespace = "system"
}

output "bgp_id" {
  value = data.f5xc_bgp.example.id
}
