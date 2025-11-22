# Example configuration for f5xc_fleet data source

data "f5xc_fleet" "example" {
  name      = "existing-fleet"
  namespace = "system"
}

output "fleet_id" {
  value = data.f5xc_fleet.example.id
}
