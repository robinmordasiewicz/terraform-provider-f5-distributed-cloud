# Example configuration for f5xc_route data source

data "f5xc_route" "example" {
  name      = "existing-route"
  namespace = "system"
}

output "route_id" {
  value = data.f5xc_route.example.id
}
